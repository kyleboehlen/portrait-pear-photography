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
	_, err = r.db.Exec(query, photo.ShootID, photo.ID)
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

type FilterPhotosParameters struct {
	ShootID    int   `json:"shoot_id"`
	Categories []int `json:"categories"`
	Favorites  bool  `json:"favorites"`
}

func (r *SQLRepo) GetPhotosFiltered(filterParams FilterPhotosParameters) ([]models.Photo, error) {
	var photos []models.Photo
	var sb strings.Builder
	args := make([]interface{}, 0)

	sb.WriteString(`
  SELECT DISTINCT p.id, p.guid, p.favorite
  FROM photos p`)

	// Join with shoots_photos only if ShootID is provided
	if filterParams.ShootID > 0 {
		sb.WriteString(`
  INNER JOIN shoots_photos sp ON p.id = sp.photo_id`)
	}

	// We need at least one WHERE clause
	sb.WriteString(`
  WHERE 1=1`)

	if filterParams.ShootID > 0 {
		sb.WriteString(` AND sp.shoot_id = ?`)
		args = append(args, filterParams.ShootID)
	}

	if filterParams.Favorites {
		sb.WriteString(` AND p.favorite = ?`)
		args = append(args, filterParams.Favorites)
	}

	if len(filterParams.Categories) > 0 {
		placeholders := make([]string, len(filterParams.Categories))
		for i, categoryID := range filterParams.Categories {
			placeholders[i] = "?"
			args = append(args, categoryID)
		}
		sb.WriteString(fmt.Sprintf(` AND p.id IN (
   SELECT pc.photo_id
   FROM photos_categories pc
   WHERE pc.category_id IN (%s)
  )`, strings.Join(placeholders, ",")))
	}

	rows, err := r.db.Query(sb.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query photos: %v", err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	photoMap := make(map[int]*models.Photo)
	var photoIDs []int

	for rows.Next() {
		var photo models.Photo
		if err := rows.Scan(&photo.ID, &photo.Guid, &photo.Favorite); err != nil {
			return nil, fmt.Errorf("failed to scan photo: %v", err)
		}

		if filterParams.ShootID > 0 {
			photo.ShootID = filterParams.ShootID
		}
		photo.Categories = []int{}

		photoMap[photo.ID] = &photo
		photoIDs = append(photoIDs, photo.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating photos: %v", err)
	}

	// Get all categories for all photos in one query
	if len(photoIDs) > 0 {
		placeholders := make([]string, len(photoIDs))
		categoryArgs := make([]interface{}, len(photoIDs))
		for i, id := range photoIDs {
			placeholders[i] = "?"
			categoryArgs[i] = id
		}

		query := fmt.Sprintf(`
   SELECT photo_id, category_id
   FROM photos_categories
   WHERE photo_id IN (%s)`, strings.Join(placeholders, ","))

		categoryRows, err := r.db.Query(query, categoryArgs...)
		if err != nil {
			return nil, fmt.Errorf("failed to query photo categories: %v", err)
		}
		defer func(categoryRows *sql.Rows) {
			_ = categoryRows.Close()
		}(categoryRows)

		for categoryRows.Next() {
			var photoID, categoryID int
			if err := categoryRows.Scan(&photoID, &categoryID); err != nil {
				return nil, fmt.Errorf("failed to scan category: %v", err)
			}
			if photo, exists := photoMap[photoID]; exists {
				photo.Categories = append(photo.Categories, categoryID)
			}
		}

		if err = categoryRows.Err(); err != nil {
			return nil, fmt.Errorf("error iterating categories: %v", err)
		}
	}

	// Convert map to slice
	for _, id := range photoIDs {
		photos = append(photos, *photoMap[id])
	}

	return photos, nil
}

func (r *SQLRepo) DeletePhoto(id int) error {
	query := `DELETE FROM photos WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete photo: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("photo with ID %d not found", id)
	}

	return nil
}
