package Services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Utils"
	geo "github.com/paulmach/go.geo"
)

// Listing represents the listing structure
type Listing struct {
	ListingID   int        `json:"listing_id"`
	Type        string     `json:"type"`     // Enum: 'Request' or 'Offer'
	Location    *geo.Point `json:"location"` // Custom handling may be needed for POINT
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DateCreated string     `json:"date_created"` // Format as "2006-01-02 15:04:05"
	Active      bool       `json:"active"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
}

// ListingService is the service layer for listing-related operations
type ListingService struct {
	db *sql.DB
}

// ValidateCoordinates validates longitude and latitude, returning city and country
func (s *ListingService) ValidateCoordinates(ctx context.Context, lat, lon float64) (string, string, error) {
	city, country, err := Utils.ReverseGeocode(lat, lon)
	if err != nil {
		return "", "", fmt.Errorf("error validating coordinates: %w", err)
	}
	return city, country, nil
}

// Reusable function to query listings based on different conditions
func (s *ListingService) queryListings(ctx context.Context, query string, args ...interface{}) ([]Listing, error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve listings: %v", err)
	}
	defer rows.Close()

	var listings []Listing
	for rows.Next() {
		var listing Listing
		if err := rows.Scan(&listing.ListingID, &listing.Type, &listing.Location, &listing.UserID, &listing.Title, &listing.Description, &listing.DateCreated, &listing.Active, &listing.City, &listing.Country); err != nil {
			return nil, fmt.Errorf("could not scan listing: %v", err)
		}
		listings = append(listings, listing)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over listings: %v", err)
	}

	// Ensure an empty slice is returned if no rows were found
	if len(listings) == 0 {
		return []Listing{}, nil
	}

	return listings, nil
}

// Create a new listing
func (s *ListingService) Create(ctx context.Context, listing *Listing) error {
	//fmt.Println(listing.Location.Lng(), listing.Location.Lat())
	city, country, err := Utils.ReverseGeocode(listing.Location.Lat(), listing.Location.Lng())
	//fmt.Println(country, city, err)
	if err != nil {
		return err
	}

	listing.City = city
	listing.Country = country
	// Convert the location to WKT format
	locationWKT := listing.Location.ToWKT()

	query := `
		INSERT INTO listings (type, location, user_id, title, description, city, country)
		VALUES (?, ST_GeomFromText(?), ?, ?, ?, ?, ?)
	`
	_, err = s.db.ExecContext(ctx, query, listing.Type, locationWKT, listing.UserID, listing.Title, listing.Description, listing.City, listing.Country)
	if err != nil {
		//fmt.Println(listing)
		return fmt.Errorf("could not create listing: %v", err)
	}

	return nil
}

// Update an existing listing
func (s *ListingService) Update(ctx context.Context, listing *Listing, listingID int) error {
	city, country, err := s.ValidateCoordinates(ctx, listing.Location.Lat(), listing.Location.Lng())
	if err != nil {
		return err
	}

	locationWKT := listing.Location.ToWKT()

	query := `
		UPDATE listings
		SET title = ?, description = ?, location = ST_GeomFromText(?), type = ?, city = ?, country = ?
		WHERE listing_id = ?
	`
	_, err = s.db.ExecContext(ctx, query, listing.Title, listing.Description, locationWKT, listing.Type, city, country, listingID)
	if err != nil {
		return fmt.Errorf("could not update listing: %v", err)
	}

	return nil
}

// Delete a listing
func (s *ListingService) Delete(ctx context.Context, listingID int) error {
	query := `DELETE FROM listings WHERE listing_id = ?`
	result, err := s.db.ExecContext(ctx, query, listingID)
	if err != nil {
		return fmt.Errorf("could not delete listing: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check affected rows: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("listing not found")
	}

	return nil
}

// Get listings by location within a radius (QueryByLocation)
func (s *ListingService) QueryByLocation(ctx context.Context, lat, lon, radius float64, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE ST_Distance(location, ST_GeomFromText('POINT(?, ?)')) < ? AND type = ?`
		return s.queryListings(ctx, query, lon, lat, radius, listingType)
	} else {
		query := `SELECT * FROM listings WHERE ST_Distance(location, ST_GeomFromText('POINT(?, ?)')) < ?`
		return s.queryListings(ctx, query, lon, lat, radius)
	}
}

// Get listings ordered by date created, descending (GetByDateCreatedDescending)
func (s *ListingService) GetByDateCreatedDescending(ctx context.Context, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE type = ? ORDER BY date_created DESC`
		return s.queryListings(ctx, query, listingType)
	} else {
		query := `SELECT * FROM listings ORDER BY date_created DESC`
		return s.queryListings(ctx, query)
	}
}

// Get listings ordered by date created and search term (GetByDateCreatedAndSearchDescending)
func (s *ListingService) GetByDateCreatedAndSearchDescending(ctx context.Context, searchTerm, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?) AND type = ? ORDER BY date_created DESC`
		return s.queryListings(ctx, query, "%"+searchTerm+"%", "%"+searchTerm+"%", listingType)
	} else {
		query := `SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?) ORDER BY date_created DESC`
		return s.queryListings(ctx, query, "%"+searchTerm+"%", "%"+searchTerm+"%")
	}
}

// Get listings by a search query in title or description (GetBySearch)
func (s *ListingService) GetBySearch(ctx context.Context, searchTerm, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?) AND type = ?`
		return s.queryListings(ctx, query, "%"+searchTerm+"%", "%"+searchTerm+"%", listingType)
	} else {
		query := `SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?)`
		return s.queryListings(ctx, query, "%"+searchTerm+"%", "%"+searchTerm+"%")
	}
}

// Get listings by distance from a specific latitude and longitude (GetByDistance)
func (s *ListingService) GetByDistance(ctx context.Context, latitude, longitude, maxDistance float64, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE ST_Distance(location, ST_GeomFromText('POINT(?, ?)')) < ? AND type = ?`
		return s.queryListings(ctx, query, longitude, latitude, maxDistance, listingType)
	} else {
		query := `SELECT * FROM listings WHERE ST_Distance(location, ST_GeomFromText('POINT(?, ?)')) < ?`
		return s.queryListings(ctx, query, longitude, latitude, maxDistance)
	}
}

// Get all listings (GetAll)
func (s *ListingService) GetAll(ctx context.Context, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE type = ?`
		return s.queryListings(ctx, query, listingType)
	} else {
		query := `SELECT * FROM listings`
		return s.queryListings(ctx, query)
	}
}

// Get listings by user ID (GetByUserID)
func (s *ListingService) GetByUserID(ctx context.Context, userID int, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE user_id = ? AND type = ?`
		return s.queryListings(ctx, query, userID, listingType)
	} else {
		query := `SELECT * FROM listings WHERE user_id = ?`
		return s.queryListings(ctx, query, userID)
	}
}

// Get a listing by its ID (GetByID)
func (s *ListingService) GetByID(ctx context.Context, listingID int) (Listing, error) {
	query := `SELECT * FROM listings WHERE listing_id = ?`
	rows, err := s.db.QueryContext(ctx, query, listingID)
	if err != nil {
		return Listing{}, fmt.Errorf("could not retrieve listing: %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		var listing Listing
		if err := rows.Scan(&listing.ListingID, &listing.Type, &listing.Location, &listing.UserID, &listing.Title, &listing.Description, &listing.DateCreated, &listing.Active, &listing.City, &listing.Country); err != nil {
			return Listing{}, fmt.Errorf("could not scan listing: %v", err)
		}
		return listing, nil
	}
	return Listing{}, fmt.Errorf("listing not found")
}
