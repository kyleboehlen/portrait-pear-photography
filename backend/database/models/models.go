package models

type Shoot struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Date string `json:"date"` // ISO 8601 format (YYYY-MM-DD)
	// TODO: default categories?
}
