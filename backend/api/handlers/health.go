package handlers

import (
	"friday/api/response"
	"friday/api/routing"
	"net/http"
)

var HealthRoutes = []routing.Route{
	{
		Method: "GET",
		Path:   "/health",
		Handle: func(w http.ResponseWriter, r *http.Request) {
			response.WriteJSONSuccessResponse(w, "Heartbeat")
		},
	},
}

func init() {
	routing.RegisterRoutes(HealthRoutes)
}
