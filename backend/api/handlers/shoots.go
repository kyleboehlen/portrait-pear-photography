package handlers

import (
	"friday/api/response"
	"friday/api/util"
	"friday/database/models"
	"friday/database/repository"
	"net/http"
)

func GetUpsertShootRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var shoot models.Shoot
		if result, err := util.ReadRequestBody(w, r, &shoot); err != nil || !result {
			return
		}

		// Create a database connection
		db := repository.Get()

		// We need an ID for updates, or a Name for creates
		if shoot.ID != 0 {
			if err := db.UpdateShoot(&shoot); err != nil {
				response.WriteJSONErrorResponse(w, "Failed to update shoot", response.ErrorCodeFailedToUpdateShoot)
				return
			}
		} else if shoot.Name != "" {
			// Create new shoot
			err := db.CreateShoot(&shoot)
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to create shoot", response.ErrorCodeFailedToCreateShoot)
				return
			}
		} else { // We need an ID for updates or a Name for creates - at least one is required
			response.WriteJSONErrorResponse(w, "Shoots requests require ID or Name", response.ErrorCodeShootsMissingRequiredFields)
			return
		}

		// This returns the updated shoot, or the created shoot with the new ID
		response.WriteJSONSuccessResponse(w, &shoot)
	}
}

func GetListShootsRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Create a database connection
		db := repository.Get()

		shoots, err := db.GetShoots()
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to get shoots", response.ErrorCodeFailedToGetShoots)
			return
		}

		response.WriteJSONSuccessResponse(w, &shoots)
	}
}

type DeleteShootRequest struct {
	ID int `json:"id"`
}

func GetDeleteShootRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deleteShootRequest DeleteShootRequest
		if result, err := util.ReadRequestBody(w, r, &deleteShootRequest); err != nil || !result {
			return
		}
		if deleteShootRequest.ID == 0 {
			response.WriteJSONErrorResponse(w, "Request is missing shoot ID", response.ErrorCodeShootsMissingRequiredFields)
			return
		}

		// Create a database connection
		db := repository.Get()

		err := db.DeleteShoot(deleteShootRequest.ID)
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to delete shoot", response.ErrorCodeFailedToDeleteShoot)
			return
		}

		response.WriteJSONSuccessResponse(w, nil)
	}
}
