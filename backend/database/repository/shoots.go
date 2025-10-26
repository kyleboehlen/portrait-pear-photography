package repository

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"friday/database/models"
	"strconv"
	"strings"
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

func (r *SQLRepo) GetShoots() ([]*models.Shoot, error) {
	query := `
  SELECT
   s.id, s.name, s.date, s.slug,
   GROUP_CONCAT(sc.category_id) as category_ids
  FROM shoots s
  LEFT JOIN shoots_categories sc ON s.id = sc.shoot_id
  GROUP BY s.id
  ORDER BY s.id
 `
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get shoots: %w", err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var shoots []*models.Shoot
	for rows.Next() {
		var shoot models.Shoot
		var categoryIDs sql.NullString

		err := rows.Scan(&shoot.ID, &shoot.Name, &shoot.Date, &shoot.Slug, &categoryIDs)
		if err != nil {
			return nil, fmt.Errorf("failed to scan shoot: %w", err)
		}

		shoot.DefaultCategories = []int{}
		if categoryIDs.Valid && categoryIDs.String != "" {
			idStrings := strings.Split(categoryIDs.String, ",")
			for _, idStr := range idStrings {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return nil, fmt.Errorf("failed to parse category ID: %w", err)
				}
				shoot.DefaultCategories = append(shoot.DefaultCategories, id)
			}
		}

		shoots = append(shoots, &shoot)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %w", err)
	}

	return shoots, nil
}
func (r *SQLRepo) DeleteShoot(id int) error {
	query := `DELETE FROM shoots WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete shoot: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("shoot with ID %d not found", id)
	}

	return nil
}

func (r *SQLRepo) UpdateShoot(shoot *models.Shoot) error {
	// Update shoot basic info (outside transaction)
	query := `UPDATE shoots SET name = ?, date = ?, slug = ? WHERE id = ?`
	result, err := r.db.Exec(query, shoot.Name, shoot.Date, shoot.Slug, shoot.ID)
	if err != nil {
		return fmt.Errorf("failed to update shoot: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("shoot with ID %d not found", shoot.ID)
	}

	// Handle categories in transaction
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	// Delete existing category associations
	query = `DELETE FROM shoots_categories WHERE shoot_id = ?`
	_, err = tx.Exec(query, shoot.ID)
	if err != nil {
		return fmt.Errorf("failed to delete shoot categories: %w", err)
	}

	// Insert new category associations
	if len(shoot.DefaultCategories) > 0 {
		insertQuery := `INSERT INTO shoots_categories (shoot_id, category_id) VALUES (?, ?)`
		for _, categoryID := range shoot.DefaultCategories {
			_, err = tx.Exec(insertQuery, shoot.ID, categoryID)
			if err != nil {
				return fmt.Errorf("failed to insert shoot category: %w", err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func generateRandomSlug(length int) string {
	bytes := make([]byte, length/2)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)[:length]
}
