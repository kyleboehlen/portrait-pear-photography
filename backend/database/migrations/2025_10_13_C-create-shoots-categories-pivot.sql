-- Create shoots_categories pivot table
CREATE TABLE IF NOT EXISTS shoots_categories (
    shoot_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    PRIMARY KEY (shoot_id, category_id),
    FOREIGN KEY (shoot_id) REFERENCES shoots(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- Create indexes for faster joins
CREATE INDEX IF NOT EXISTS idx_shoots_categories_shoot_id ON shoots_categories(shoot_id);
CREATE INDEX IF NOT EXISTS idx_shoots_categories_category_id ON shoots_categories(category_id);