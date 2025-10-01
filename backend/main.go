package main

import (
	_ "friday/api/handlers" // This is required to register the routes
	"friday/api/middleware"
	"friday/api/routing"
	"log"
	"net/http"
)

func main() {
	// This gets a new server mux with all routes registered, this allows routes to be defined in their own file
	// alongside the handlers they're associated with
	mux := routing.GetNewRouter()

	// Reflect Origin to satisfy CORS requirements for pre-flight checks - we may need to limit this in the future to
	// origins specified in the .env vars
	httpHandler := middleware.CORS(mux)

	// Health - need to add db status checks

	// Get all home page photos
	// Get photos from shoot

	// Admin - Upsert shoot
	// Admin - Delete shoot
	// Admin - Upsert photo
	// Admin - Delete photo

	// Admin - set home page photos

	// Route groups
	// handlers map one-to-one functions with routes
	// Handlers call services, they orchestrate between basic functions, like deleting a shoot might call delete photos first then delete shoot
	// Repositories interact with outside services, database, file storage, cloudflare, etc.

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal(err)
	}
}
