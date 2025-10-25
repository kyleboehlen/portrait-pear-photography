package models

type Shoot struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Date              string `json:"date"` // ISO 8601 format (YYYY-MM-DD)
	DefaultCategories []int  `json:"default_categories"`
}

type Photo struct {
	ID         int    `json:"id"`
	Guid       string `json:"guid"`
	Favorite   bool   `json:"favorite"`
	ShootID    int    `json:"shoot_id"`
	Categories []int  `json:"categories"`
}
