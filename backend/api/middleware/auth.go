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
			if valid, err := auth.IsJWTValid(token); err != nil || !valid {
				response.WriteJSONErrorResponse(w, "Invalid or expired token", response.ErrorCodeInvalidOrExpiredToken)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
