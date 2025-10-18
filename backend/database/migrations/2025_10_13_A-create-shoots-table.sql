-- Create shoots table
CREATE TABLE IF NOT EXISTS shoots
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT NOT NULL,
    date       DATE,
    -- NOTE: Uniqueness here needs to be enforced at the application level
    slug       VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster lookups by slug
CREATE INDEX IF NOT EXISTS idx_shoots_name ON shoots (name);
CREATE INDEX IF NOT EXISTS idx_shoots_slug ON shoots (slug);