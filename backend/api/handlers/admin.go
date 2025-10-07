package handlers

import (
	"encoding/json"
	"friday/api/response"
	"friday/api/routing"
	"friday/services/auth"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type AdminAuthRequest struct {
	Password string `json:"password"`
}

type AdminAuthResponse struct {
	Token string `json:"token"`
}

var AdminRoutes = []routing.Route{
	{
		Method: "POST",
		Path:   "/authenticate", // Don't use admin prefix or AuthN middleware will block this route
		Handle: func(w http.ResponseWriter, r *http.Request) {
			// TODO: this needs to be moved to a func that can be passed in a struct for a request, and return the unmarshaled body. We'll for sure reuse this

			body, err := io.ReadAll(r.Body)
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to read request body", response.ErrorCodeRequestBodyUnreadable)
				return
			}
			defer r.Body.Close()

			var adminAuthReq AdminAuthRequest
			if err := json.Unmarshal(body, &adminAuthReq); err != nil {
				response.WriteJSONErrorResponse(w, "Failed to unmarshal request body", response.ErrorCodeRequestJSONUnmarshalFailed)
				return
			}

			// TODO: Replace hardcoded hash
			if err := bcrypt.CompareHashAndPassword([]byte("$2y$15$9TqhPk5H89hguNdvoJIUA.27.MjDh6lDlRo/n.zBiiAPnkKXkK//W"), []byte(adminAuthReq.Password)); err != nil {
				response.WriteJSONErrorResponse(w, "Invalid password", response.ErrorCodeInvalidPassword)
				return
			}

			// TODO: Remove the hardcoded secret
			token, err := auth.CreateAdminJWT("861a2703c15c319bd379d9a776232872ef5b82968a4b65627cf4d689def3b9383647849df7899ae80bd9cf0f8c7b1edb03472f6596498648a299cbf36aad837b")
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to create JWT", response.ErrorCodeJWTGenerationFailed)
				return
			}
			response.WriteJSONSuccessResponse(w, AdminAuthResponse{Token: token})
		},
	},
}

func init() {
	routing.RegisterRoutes(AdminRoutes)
}
