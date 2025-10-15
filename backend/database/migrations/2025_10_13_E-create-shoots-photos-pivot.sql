-- Create shoots_photos pivot table
CREATE TABLE IF NOT EXISTS shoots_photos (
    shoot_id INTEGER NOT NULL,
    photo_id INTEGER NOT NULL,
    PRIMARY KEY (shoot_id, photo_id),
    FOREIGN KEY (shoot_id) REFERENCES shoots(id) ON DELETE CASCADE,
    FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE
);

-- Create indexes for faster joins
CREATE INDEX IF NOT EXISTS idx_shoots_photos_shoot_id ON shoots_photos(shoot_id);
CREATE INDEX IF NOT EXISTS idx_shoots_photos_photo_id ON shoots_photos(photo_id);