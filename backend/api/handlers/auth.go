package handlers

import (
	"friday/api/response"
	"friday/api/util"
	"friday/services/auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AdminAuthRequest struct {
	Password string `json:"password"`
}

type AdminAuthResponse struct {
	Token string `json:"token"`
}

func GetAuthenticateRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var adminAuthReq AdminAuthRequest
		if ok, err := util.ReadRequestBody(w, r, &adminAuthReq); err != nil || !ok {
			return
		}
		if adminAuthReq.Password == "" {
			response.WriteJSONErrorResponse(w, "Request is missing password", response.ErrorCodeMissingAdminPassword)
			return
		}

		passwordHash, exists := auth.GetCurrentHash()
		if !exists {
			response.WriteJSONErrorResponse(w, "Admin password hash not set", response.ErrorCodeAdminPasswordNotSet)
			return
		}

		now := time.Now()
		salt := now.Format("2006-01")
		if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(adminAuthReq.Password+salt)); err != nil {
			response.WriteJSONErrorResponse(w, "Invalid password", response.ErrorCodeInvalidPassword)
			return
		}

		token, err := auth.CreateAdminJWT()
		if err != nil {
			response.WriteJSONErrorResponse(w, "Failed to create JWT", response.ErrorCodeJWTGenerationFailed)
			return
		}

		response.WriteJSONSuccessResponse(w, AdminAuthResponse{Token: token})
	}
}

type AdminTestResponse struct {
	Test string `json:"test"`
}

func GetTestRouteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJSONSuccessResponse(w, &AdminTestResponse{Test: "asdf"})
	}
}
