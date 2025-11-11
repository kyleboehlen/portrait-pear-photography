package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"friday/api/response"
	"friday/api/util"
	"friday/database/models"
	"friday/database/repository"
	"friday/services/images"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
)

type UploadPhotoExtraDataRequest struct {
	ShootID    int   `json:"shoot_id"`
	Categories []int `json:"categories"`
}

const (
	MaxFileSize = 20 * 1024 * 1024 // 20MB
	TargetSize  = 19 * 1024 * 1024 // 19MB
)

// CompressedFile implements multipart.File interface
type CompressedFile struct {
	*bytes.Reader
	size int64
}

func (cf *CompressedFile) Close() error {
	// bytes.Reader doesn't need closing, but we implement it for the interface
	return nil
}

func NewCompressedFile(data []byte) *CompressedFile {
	return &CompressedFile{
		Reader: bytes.NewReader(data),
		size:   int64(len(data)),
	}
}

func (cf *CompressedFile) Size() int64 {
	return cf.size
}

func GetUploadPhotoRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Limit upload size (50 MB)
		r.Body = http.MaxBytesReader(w, r.Body, 50<<20)
		if err := r.ParseMultipartForm(50 << 20); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to parse photo form data", response.ErrorCodeFailedToReadUploadPhotoRequest)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			response.WriteJSONErrorResponse(w, "Request is missing the photo file", response.ErrorCodeMissingPhotoFileRequest)
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

		// Check file size
		fileSize := header.Size

		var smallEnoughForUploadFile *CompressedFile

		fileBytes, err := io.ReadAll(file)
		if fileSize <= MaxFileSize {
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to read image bytes", response.ErrorCodeFailedToCompressImage)
			}
			smallEnoughForUploadFile = NewCompressedFile(fileBytes)
		} else {
			// File is > 20MB, need to compress
			compressedFile, err := compressJPEGHighQuality(fileBytes, TargetSize)
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to compress image", response.ErrorCodeFailedToCompressImage)
				return
			}
			smallEnoughForUploadFile = compressedFile
		}

		// Need to make sure that we can get the extra info for the photo's shoot and categories
		extraDataJSON := r.FormValue("data")
		if extraDataJSON == "" {
			response.WriteJSONErrorResponse(w, "Missing extra photo data", response.ErrorCodeMissingExtraPhotoData)
			return
		}

		var extraData UploadPhotoExtraDataRequest
		if err := json.Unmarshal([]byte(extraDataJSON), &extraData); err != nil {
			response.WriteJSONErrorResponse(w, "Invalid extra data format", response.ErrorCodeInvalidExtraPhotoData)
			return
		}

		cloudflareClient := images.Get()
		image, err := cloudflareClient.UploadPhoto(smallEnoughForUploadFile)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to upload photo: "+err.Error(), response.ErrorCodeFailedToUploadToCloudflare)
			return
		}

		var photo = &models.Photo{
			Guid:       image.ID,
			ShootID:    extraData.ShootID,
			Categories: extraData.Categories,
		}

		// Insert the photo in to the database
		db := repository.Get()
		err = db.CreatePhoto(photo)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to save photo to database", response.ErrorCodeFailedToCreatePhoto)
			return
		}

		response.WriteJSONSuccessResponse(w, photo)
	}
}

func compressJPEGHighQuality(fileBytes []byte, targetSize int64) (*CompressedFile, error) {

	// Decode JPEG image
	img, err := jpeg.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode JPEG: %w", err)
	}

	// Try different quality levels starting from high quality
	for quality := 95; quality >= 10; quality -= 5 {
		var buf bytes.Buffer

		opts := &jpeg.Options{Quality: quality}
		err = jpeg.Encode(&buf, img, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to reduce JPEG to target size: %w", err)
		}

		// Check if we've reached target size
		if int64(buf.Len()) <= targetSize {
			return NewCompressedFile(buf.Bytes()), nil
		}
	}

	return nil, errors.New("unable to compress image to target size")
}

type DeletePhotoRequest struct {
	ID int `json:"id"`
}

func GetDeletePhotoRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deletePhotoRequest DeletePhotoRequest
		if result, err := util.ReadRequestBody(w, r, &deletePhotoRequest); err != nil || !result {
			return
		}
		if deletePhotoRequest.ID == 0 {
			response.WriteJSONErrorResponse(w, "Request is missing photo ID", response.ErrorCodePhotosMissingRequiredFields)
			return
		}

		// Create a database connection
		db := repository.Get()

		err := db.DeletePhoto(deletePhotoRequest.ID)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to delete photo", response.ErrorCodeFailedToDeletePhoto)
			return
		}

		response.WriteJSONSuccessResponse(w, nil)
	}
}

type FilterPhotosRequest struct {
	FilterParameters repository.FilterPhotosParameters `json:"filter_parameters"`
}

func GetListPhotosFilteredRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var filterPhotosRequest FilterPhotosRequest
		if result, err := util.ReadRequestBody(w, r, &filterPhotosRequest); err != nil || !result {
			return
		}

		db := repository.Get()
		photos, err := db.GetPhotosFiltered(filterPhotosRequest.FilterParameters)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to get photos", response.ErrorCodeFailedToGetPhotos)
			return
		}

		response.WriteJSONSuccessResponse(w, &photos)
	}
}

type ShootSlugRequest struct {
	ShootSlug string `json:"shoot_slug"`
}

func GetListShootPhotosRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var shootSlugRequest ShootSlugRequest
		if result, err := util.ReadRequestBody(w, r, &shootSlugRequest); err != nil || !result {
			return
		}

		db := repository.Get()
		shoot, err := db.GetShootBySlug(shootSlugRequest.ShootSlug)
		photos, err := db.GetPhotosFiltered(repository.FilterPhotosParameters{
			ShootID: shoot.ID,
		})
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to get photos", response.ErrorCodeFailedToGetPhotos)
			return
		}

		response.WriteJSONSuccessResponse(w, &photos)
	}
}

func GetListFavoritePhotosRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := repository.Get()

		photos, err := db.GetPhotosFiltered(repository.FilterPhotosParameters{
			Favorites: true,
		})
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to get favorite photos", response.ErrorCodeFailedToGetPhotos)
			return
		}

		response.WriteJSONSuccessResponse(w, &photos)
	}
}

type UpdatePhotoRequest struct {
	Photo models.Photo `json:"photo"`
}

func GetUpdatePhotoRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updatePhotoRequest UpdatePhotoRequest
		if result, err := util.ReadRequestBody(w, r, &updatePhotoRequest); err != nil || !result {
			return
		}

		db := repository.Get()
		err := db.UpdatePhoto(&updatePhotoRequest.Photo)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to update photo", response.ErrorCodeFailedToUpdatePhoto)
			return
		}

		response.WriteJSONSuccessResponse(w, &updatePhotoRequest.Photo)
	}
}
