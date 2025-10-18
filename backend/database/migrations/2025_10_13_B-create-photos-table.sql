-- Create photos table
CREATE TABLE IF NOT EXISTS photos
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    guid     TEXT    NOT NULL UNIQUE,
    favorite BOOLEAN NOT NULL DEFAULT 0
);

-- Create index for faster lookups by guid and favorite status
CREATE INDEX IF NOT EXISTS idx_photos_guid ON photos (guid);
CREATE INDEX IF NOT EXISTS idx_photos_favorite ON photos (favorite);