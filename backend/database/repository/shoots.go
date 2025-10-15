package repository

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"friday/database/models"
	"time"
)

func (r *SQLRepo) CreateShoot(shoot *models.Shoot) error {
	query := `INSERT INTO shoots (name, date, slug, created_at) VALUES (?, ?, ?, datetime('now'))`

	// Needed to satisfy the unique constraint
	if shoot.Slug == "" {
		shoot.Slug = generateRandomSlug(16)
	}

	// Default to today's date to provide easy update in the next frontend workflow step
	if shoot.Date == "" {
		shoot.Date = time.Now().Format("2006-01-02")
	}

	result, err := r.db.Exec(query, shoot.Name, shoot.Date, shoot.Slug)
	if err != nil {
		return fmt.Errorf("failed to create shoot: %w", err)
	}

	// Get the auto-generated ID and set it on the Shoot
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get latest shoot ID: %w", err)
	}

	shoot.ID = int(id)
	return nil
}

func generateRandomSlug(length int) string {
	bytes := make([]byte, length/2)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)[:length]
}
