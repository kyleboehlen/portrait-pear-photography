-- Create migrations table to track applied database migrations
CREATE TABLE IF NOT EXISTS migrations
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    filename   TEXT NOT NULL UNIQUE,
    applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster lookups by filename
CREATE INDEX IF NOT EXISTS idx_migrations_filename ON migrations (filename);