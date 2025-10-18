-- Create photos_categories pivot table
CREATE TABLE IF NOT EXISTS photos_categories
(
    photo_id    INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    PRIMARY KEY (photo_id, category_id),
    FOREIGN KEY (photo_id) REFERENCES photos (id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);

-- Create indexes for faster joins
CREATE INDEX IF NOT EXISTS idx_photos_categories_photo_id ON photos_categories (photo_id);
CREATE INDEX IF NOT EXISTS idx_photos_categories_category_id ON photos_categories (category_id);