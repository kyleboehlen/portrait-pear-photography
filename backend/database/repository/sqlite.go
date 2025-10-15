package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLRepo struct {
	db *sql.DB
}

func Setup() (*SQLRepo, error) {
	return SetupWithMigration(false)
}

// SetupWithMigration Responsible for creating/checking for the database file and running migrations. This also returns the SQLRepo struct ^^^
func SetupWithMigration(migrate bool) (*SQLRepo, error) {
	dbDir := "database/data"

	// Ensure the database directory exists, the directory will always exist locally with the .gitignore placeholder,
	// however the deployed binary will not include it
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	// Check if database file exists
	// _, err := os.Stat(dbPath)
	//isNewDB := os.IsNotExist(err) // Placeholder for new database logic if needed in the future

	// Database file path
	dbPath := filepath.Join(dbDir, "friday.db")

	// SQLite will create the database file on open if it doesn't exist
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set value to struct
	repo := &SQLRepo{db: db}

	// Only run the file validation/migration if migrate is true, this is done in main.go on startup
	if migrate {
		_ = repo.runMigrations()
	}

	return repo, nil
}

func (r *SQLRepo) PingDB() (bool, error) {
	if err := r.db.Ping(); err != nil {
		return false, fmt.Errorf("failed to ping database: %w", err)
	}
	return true, nil
}

// Method for running migrations, uses the repo
func (r *SQLRepo) runMigrations() error {
	migrationDir := "database/migrations"

	// Get migration files
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Get applied migrations
	appliedMigrations, err := r.getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Run new migrations in order
	for _, file := range files {
		// There should NEVER be a
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		if appliedMigrations[filename] {
			continue
		}

		// Read and execute migration
		content, err := os.ReadFile(filepath.Join(migrationDir, filename))
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

func (r *SQLRepo) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}
