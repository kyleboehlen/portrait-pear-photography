package main

import (
	_ "friday/api/handlers" // This is required to register the routes
	"friday/api/middleware"
	"friday/api/routing"
	"friday/database/repository"
	"log"
	"net/http"
)

func main() {
	// This gets a new server mux with all routes registered, this allows routes to be defined in their own file
	// alongside the handlers they're associated with
	mux := routing.GetNewRouter()

	// NOTE: Middleware steps will be executed in reverse order of how they are applied here

	// Check auth for admin routes
	httpHandler := middleware.AuthenticateAdmin(mux)

	// Reflect Origin to satisfy CORS requirements for pre-flight checks - we may need to limit this in the future to
	// origins specified in the .env vars
	httpHandler = middleware.CORS(httpHandler)

	// Setup database - run migrations and check for database file
	_, err := repository.SetupWithMigration(true)
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}

	// Get all home page photos
	// Get photos from shoot
	// Get all shoots

	// Admin - Delete shoot
	// Admin - Upsert photo
	// Admin - Delete photo

	// Admin - set home page photos

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal(err)
	}
}
