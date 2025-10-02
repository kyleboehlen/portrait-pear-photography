package handlers

import (
	"friday/api/response"
	"friday/api/routing"
	"friday/database/repository"
	"net/http"
)

var HealthRoutes = []routing.Route{
	{
		Method: "GET",
		Path:   "/health",
		Handle: func(w http.ResponseWriter, r *http.Request) {
			repo, err := repository.Setup()
			if err != nil {
				response.WriteJSONErrorResponse(w, "Failed to setup DB repo", response.ErrorCodeDatabaseConnection)
			}

			if success, err := repo.PingDB(); err != nil || !success {
				response.WriteJSONErrorResponse(w, "Database ping failed", response.ErrorCodeDatabaseNotHealthy)
				return
			}

			// TODO: Check the health of our Cloudflare Images connection

			response.WriteJSONSuccessResponse(w, "Heartbeat")
		},
	},
}

func init() {
	routing.RegisterRoutes(HealthRoutes)
}
