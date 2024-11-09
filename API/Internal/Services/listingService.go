package Services

import (
	"database/sql"
	"time"
)

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Listing struct {
	ListingID   int       `json:"listing_id"`
	Type        string    `json:"type"`     // Enum: 'Request' or 'Offer'
	Location    Point     `json:"location"` // Custom handling may be needed for POINT
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"` // Format as "2006-01-02 15:04:05"
	Active      bool      `json:"active"`
}

type ListingService struct {
	db *sql.DB
}
