package repository

import (
	"database/sql"
	"fmt"
	"friday/database/models"
	"strings"
)

func (r *SQLRepo) CreatePhoto(photo *models.Photo) error {
	query := `INSERT INTO photos (guid, favorite) VALUES (?, false)`

	result, err := r.db.Exec(query, photo.Guid, photo.Favorite)
	if err != nil {
		return fmt.Errorf("failed to create photo: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get latest photo ID: %v", err)
	}

	photo.ID = int(id)

	// Set shoot ID - photos can only be uploaded in the scope of a shoot, this doesn't need to be editable
	query = `INSERT INTO shoots_photos (shoot_id, photo_id) VALUES (?, ?)`
	_, err = r.db.Exec(query, photo.ID, photo.ShootID)
	if err != nil {
		return fmt.Errorf("failed to set shoot for the photo: %v", err)
	}

	// Set photo categories and favorite status
	err = r.UpdatePhoto(photo)
	if err != nil {
		return err
	}

	// Set photo shoot ID
	return nil
}

func (r *SQLRepo) UpdatePhoto(photo *models.Photo) error {
	// Update favorite status - we know the photo row already exists
	query := `UPDATE photos SET favorite = ? WHERE id = ?`
	_, err := r.db.Exec(query, photo.Favorite, photo.ID)
	if err != nil {
		return fmt.Errorf("failed to update photo favorite status: %v", err)
	}

	// Start transaction for category updates
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	// Delete existing category associations
	query = `DELETE FROM photos_categories WHERE photo_id = ?`
	_, err = tx.Exec(query, photo.ID)
	if err != nil {
		return fmt.Errorf("failed to delete existing categories for photo: %v", err)
	}

	// Batch insert new category associations
	if len(photo.Categories) > 0 {
		valueStrings := make([]string, 0, len(photo.Categories))
		valueArgs := make([]interface{}, 0, len(photo.Categories)*2)

		for _, categoryID := range photo.Categories {
			valueStrings = append(valueStrings, "(?, ?)")
			valueArgs = append(valueArgs, photo.ID, categoryID)
		}

		var sb strings.Builder
		sb.WriteString("INSERT INTO photos_categories (photo_id, category_id) VALUES ")
		sb.WriteString(strings.Join(valueStrings, ","))
		query = sb.String()

		_, err = tx.Exec(query, valueArgs...)
		if err != nil {
			return fmt.Errorf("failed to add categories for photo: %v", err)
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
