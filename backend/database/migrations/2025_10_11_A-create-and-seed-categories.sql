-- Create categories table
CREATE TABLE IF NOT EXISTS categories
(
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Insert default categories
INSERT OR IGNORE INTO categories (id, name)
VALUES (1, 'portrait'),
       (2, 'automotive'),
       (3, 'street'),
       (4, 'blackandwhite');

-- Create index for faster lookups by name
CREATE INDEX IF NOT EXISTS idx_categories_name ON categories (name);