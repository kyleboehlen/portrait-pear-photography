package images

import (
	"context"
	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/images"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"mime/multipart"
	"os"
	"slices"
)

type CloudflareClient struct {
	client    *cloudflare.Client
	accountID string
}

var instance *CloudflareClient

func Get() *CloudflareClient {
	if instance == nil {
		instance = &CloudflareClient{}
		instance.init()
	}

	return instance
}

func (c *CloudflareClient) init() {
	client := cloudflare.NewClient(
		option.WithAPIToken(os.Getenv("CLOUDFLARE_IMAGES_API_TOKEN")))
	c.client = client
	c.accountID = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
}

func (c *CloudflareClient) UploadPhoto(photo multipart.File) (*images.Image, error) {
	image, err := c.client.Images.V1.New(context.TODO(), images.V1NewParams{
		AccountID:         cloudflare.F(c.accountID),
		File:              cloudflare.F[interface{}](photo),
		RequireSignedURLs: cloudflare.F(false),
	})
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (c *CloudflareClient) PurgePhotos(guids []string) error {
	page, err := c.client.Images.V2.List(context.Background(), images.V2ListParams{
		AccountID: cloudflare.F(c.accountID),
		PerPage:   cloudflare.F(float64(100)),
	})

	if err != nil {
		return err
	}

	idsToDelete := make([]string, 0)
	for len(page.Images) > 0 {
		for _, image := range page.Images {
			if !slices.Contains(guids, image.ID) {
				idsToDelete = append(idsToDelete, image.ID)
			}
		}

		if page.ContinuationToken == "" {
			break
		}

		page, err = c.client.Images.V2.List(context.Background(), images.V2ListParams{
			AccountID:         cloudflare.F(c.accountID),
			ContinuationToken: cloudflare.F(page.ContinuationToken),
			PerPage:           cloudflare.F(float64(100)),
		})

		if err != nil {
			return err
		}
	}

	for _, id := range idsToDelete {
		_, err = c.client.Images.V1.Delete(
			context.TODO(),
			id,
			images.V1DeleteParams{
				AccountID: cloudflare.F(c.accountID),
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}
