package repository

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var dbDir = "database/data"
var dbFileName = "friday.db"

type SQLRepo struct {
	db *sql.DB
}

var instance *SQLRepo

// Get returns the singleton repository instance
func Get() *SQLRepo {
	if instance == nil {
		panic("repository not initialized - call Setup() first")
	} else {
		return instance
	}
}

func Setup(migrationFS embed.FS) error {
	if os.Getenv("TEST") == "true" && os.Getenv("IN_MEMORY_DATABASE") != "false" {
		repo, err := setupTestDB()
		if err != nil {
			return err
		}
		instance = repo
	} else {
		repo, err := SetupWithMigration(migrationFS)
		if err != nil {
			return err
		}
		instance = repo
	}

	return nil
}

func setupTestDB() (*SQLRepo, error) {
	// Use in-memory SQLite for tests
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	repo := &SQLRepo{db: db}

	// Run migrations/create tables
	_ = repo.runMigrationsFromDisk()

	return repo, nil
}

// SetupWithMigration Responsible for creating/checking for the database file and running migrations. This also returns the SQLRepo struct ^^^
func SetupWithMigration(migrationFS embed.FS) (*SQLRepo, error) {
	// Ensure the database directory exists, the directory will always exist locally with the .gitignore placeholder,
	// however the deployed binary will not include it
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Check if database file exists
	// _, err := os.Stat(dbPath)
	//isNewDB := os.IsNotExist(err) // Placeholder for new database logic if needed in the future

	// Database file path
	dbPath := GetDBPath()

	// SQLite will create the database file on open if it doesn't exist
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set value to struct
	repo := &SQLRepo{db: db}

	// Run the migrations with the embedded migrations filesystem from main.go
	_ = repo.runMigrations(migrationFS)

	return repo, nil
}

func (r *SQLRepo) PingDB() (bool, error) {
	if err := r.db.Ping(); err != nil {
		return false, fmt.Errorf("failed to ping database: %w", err)
	}
	return true, nil
}

// Method for running migrations, uses the repo
func (r *SQLRepo) runMigrations(migrationFS embed.FS) error {
	files, err := migrationFS.ReadDir("database/migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Get applied migrations
	appliedMigrations, err := r.getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if appliedMigrations[filename] {
			continue
		}

		content, err := migrationFS.ReadFile(filepath.Join("database/migrations", filename))
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", filename, err)
		}

		if _, err := r.db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		// Record migration
		_, err = r.db.Exec("INSERT INTO migrations (filename) VALUES (?)", filename)
		if err != nil {
			return fmt.Errorf("failed to record migration %s: %w", filename, err)
		}
	}

	return nil
}

func (r *SQLRepo) runMigrationsFromDisk() error {
	// Find the project root by looking for go.mod
	projectRoot, err := findProjectRoot()
	if err != nil {
		return fmt.Errorf("failed to find project root: %w", err)
	}

	migrationsPath := filepath.Join(projectRoot, "database", "migrations")

	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	appliedMigrations, err := r.getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if appliedMigrations[filename] {
			continue
		}

		content, err := os.ReadFile(filepath.Join(migrationsPath, filename))
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", filename, err)
		}

		if _, err := r.db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		_, err = r.db.Exec("INSERT INTO migrations (filename) VALUES (?)", filename)
		if err != nil {
			return fmt.Errorf("failed to record migration %s: %w", filename, err)
		}
	}

	return nil
}

// findProjectRoot walks up the directory tree to find go.mod
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find go.mod in parent directories")
		}
		dir = parent
	}
}

func (r *SQLRepo) getAppliedMigrations() (map[string]bool, error) {
	applied := make(map[string]bool)

	rows, err := r.db.Query("SELECT filename FROM migrations")
	if err != nil {
		// Table doesn't exist yet, return empty map
		return applied, nil
	}

	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			// Log the error or handle as needed
			fmt.Printf("Warning: failed to close rows: %v\n", closeErr)
		}
	}()

	for rows.Next() {
		var filename string
		if err := rows.Scan(&filename); err != nil {
			return nil, err
		}
		applied[filename] = true
	}

	return applied, rows.Err()
}

func (r *SQLRepo) ReloadFromDisk() error {
	dbPath := GetDBPath()

	// Open new DB
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Close existing DB (if any) via the receiver
	if err := r.Close(); err != nil {
		// If close fails, close the newly opened db to avoid leaks and return the error
		_ = db.Close()
		return fmt.Errorf("failed to close existing database: %w", err)
	}

	// Replace repo DB and return
	r.db = db
	return nil
}

func (r *SQLRepo) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

func GetDBPath() string {
	return filepath.Join(dbDir, dbFileName)
}
