package handlers

import (
	"encoding/json"
	"friday/api/response"
	"friday/api/util"
	"friday/database/models"
	"friday/database/repository"
	"friday/services/images"
	"mime/multipart"
	"net/http"
)

type UploadPhotoExtraDataRequest struct {
	ShootID    int   `json:"shoot_id"`
	Categories []int `json:"categories"`
}

func GetUploadPhotoRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Limit upload size (20 MB)
		r.Body = http.MaxBytesReader(w, r.Body, 20<<20)
		if err := r.ParseMultipartForm(20 << 20); err != nil {
			response.WriteJSONErrorResponse(w, "Failed to parse photo form data", response.ErrorCodeFailedToReadUploadPhotoRequest)
			return
		}

		// All we need is a Reader to pass to the Cloudflare API
		file, _, err := r.FormFile("file")
		if err != nil {
			response.WriteJSONErrorResponse(w, "Request is missing the photo file", response.ErrorCodeMissingPhotoFileRequest)
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

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
		image, err := cloudflareClient.UploadPhoto(file)
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
		}

		response.WriteJSONSuccessResponse(w, &photos)
	}
}
