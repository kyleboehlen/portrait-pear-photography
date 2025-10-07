package middleware

import (
	"friday/api/response"
	"friday/services/auth"
	"net/http"
	"strings"
)

// AuthenticateAdmin validates a bearer token from the Authorization header
func AuthenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We only care about admin routes, other routes are public (with CORS)
		if strings.HasPrefix(r.URL.Path, "/admin") {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.WriteJSONErrorResponse(w, "Authorization header required", response.ErrorCodeMissingAuthorizationHeader)
				return
			}

			// We're expecting a JWT in Bearer format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				response.WriteJSONErrorResponse(w, "Invalid authorization header format", response.ErrorCodeInvalidAuthorizationHeaderFormat)
				return
			}

			token := parts[1]

			// We don't actually use any claims, so we just need to know that it will parse correctly without errors
			// TODO: Remove the hardcoded secret
			if _, err := auth.ParseAdminJWT(token, "861a2703c15c319bd379d9a776232872ef5b82968a4b65627cf4d689def3b9383647849df7899ae80bd9cf0f8c7b1edb03472f6596498648a299cbf36aad837b"); err != nil {
				response.WriteJSONErrorResponse(w, "Invalid or expired token", response.ErrorCodeInvalidOrExpiredToken)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
