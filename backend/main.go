package main

import (
	"embed"
	_ "friday/api/handlers" // This is required to register the routes
	"friday/api/middleware"
	"friday/api/routing"
	"friday/database/repository"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// Need to embed the database migrations in the binary
//
//go:embed database/migrations
var migrationFS embed.FS

func main() {
	// load environment variables from .env (falls back to system env)
	_ = godotenv.Load(".env")

	// This gets a new server mux with all routes registered, this allows routes to be defined in their own file
	// alongside the handlers they're associated with
	mux := routing.GetNewRouter()

	// NOTE: Middleware steps will be executed in reverse order of how they are applied here

	// Check auth for admin routes
	httpHandler := middleware.AuthenticateAdmin(mux)

	// Reflect Origin to satisfy CORS requirements for pre-flight checks - we may need to limit this in the future to
	// origins specified in the .env vars
	httpHandler = middleware.CORS(httpHandler)

	// Setup database - run migrations using the embedded migration files, and check for the database file
	err := repository.Setup(migrationFS)
	if err != nil {
		log.Fatalf("Failed to set up database: %v", err)
	}

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", httpHandler); err != nil {
		log.Fatal(err)
	}
}
