package Services

import "database/sql"

type Image struct {
	ImageID       int    `json:"image_id"`
	URL           string `json:"url"`
	UserID        int    `json:"user_id"`
	ListingID     int    `json:"listing_id,omitempty"` // Nullable
	ShowOnProfile bool   `json:"show_on_profile"`
}

type ImageService struct {
	db *sql.DB
}
