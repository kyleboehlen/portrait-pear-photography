package images

import (
	"context"
	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/images"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"mime/multipart"
	"os"
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
