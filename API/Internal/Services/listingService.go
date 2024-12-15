package Services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Utils"
	geo "github.com/paulmach/go.geo"
)

// Listing represents a listing in the marketplace
// @Description A listing that can either be a request or offer, containing details about the title, description, and user information.
type Listing struct {
	// ListingID is the unique identifier for the listing
	// @example 101
	ListingID int `json:"listing_id"`

	// Type specifies whether the listing is a request or offer
	// Enum: 'Request', 'Offer'
	// @example "Offer"
	Type string `json:"type"`

	// Location is the geographical location of the listing
	// Custom handling may be needed for geo.Point type
	// @example {"lat": 34.0522, "lng": -118.2437}
	Location *geo.Point `json:"location"`

	// UserID is the ID of the user who created the listing
	// @example 1
	UserID int `json:"user_id"`

	// Title is the title of the listing
	// @example "Looking for a plumber"
	Title string `json:"title"`

	// Description is the detailed description of the listing
	// @example "Need a plumber for a quick job fixing a leaky pipe."
	Description string `json:"description"`

	// DateCreated is the date when the listing was created
	// Format: "2006-01-02 15:04:05"
	// @example "2024-12-16 14:30:00"
	DateCreated string `json:"date_created"`

	// Active specifies whether the listing is currently active
	// @example true
	Active bool `json:"active"`

	// City is the city where the listing is located
	// @example "Los Angeles"
	City string `json:"city"`

	// Country is the country where the listing is located
	// @example "USA"
	Country string `json:"country"`
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
		if err := rows.Scan(&listing.ListingID, &listing.Type, &listing.Location, &listing.UserID,
			&listing.Title, &listing.Description, &listing.DateCreated, &listing.Active,
			&listing.City, &listing.Country); err != nil {
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
func (s *ListingService) Create(ctx context.Context, listing *Listing) (Listing, error) {
	city, country, err := Utils.ReverseGeocode(listing.Location.Lat(), listing.Location.Lng())
	if err != nil {
		return Listing{}, fmt.Errorf("could not validate coordinates: %w", err)
	}

	listing.City = city
	listing.Country = country

	// Convert the location to WKT format
	locationWKT := listing.Location.ToWKT()

	query := `
        INSERT INTO listings (type, location, user_id, title, description, city, country)
        VALUES (?, ST_GeomFromText(?), ?, ?, ?, ?, ?)
    `
	result, err := s.db.ExecContext(ctx, query, listing.Type, locationWKT, listing.UserID, listing.Title, listing.Description, listing.City, listing.Country)
	if err != nil {
		return Listing{}, fmt.Errorf("could not create listing: %v", err)
	}

	// Get the ID of the newly inserted listing
	listingID, err := result.LastInsertId()
	if err != nil {
		return Listing{}, fmt.Errorf("could not retrieve last insert ID: %v", err)
	}

	// Retrieve the full listing by its ID
	createdListing, err := s.GetByID(ctx, int(listingID))
	if err != nil {
		return Listing{}, fmt.Errorf("could not retrieve newly created listing: %v", err)
	}

	return createdListing, nil
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

// TODO rpleace with query by city
// Get listings by location within a radius (QueryByLocation)
/*func (s *ListingService) QueryByLocation(ctx context.Context, lat, lon, radius float64, listingType string) ([]Listing, error) {
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE ST_Distance(location, ST_GeomFromText('POINT(?, ?)')) < ? AND type = ?`
		return s.queryListings(ctx, query, lon, lat, radius, listingType)
	} else {
		query := `SELECT * FROM listings WHERE ST_Distance(location, ST_GeomFromText('POINT(?, ?)')) < ?`
		return s.queryListings(ctx, query, lon, lat, radius)
	}
}*/

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

	maxDistance *= 1000
	fmt.Println(longitude, latitude, maxDistance)
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings WHERE ST_Distance_Sphere(location, ST_GeomFromText(CONCAT('POINT(', ?, ' ', ?, ')'))) < ? AND type = ?`
		return s.queryListings(ctx, query, longitude, latitude, maxDistance, listingType)
	} else {
		query := `SELECT * FROM listings WHERE ST_Distance_Sphere(location, ST_GeomFromText(CONCAT('POINT(', ?, ' ', ?, ')'))) < ?`
		return s.queryListings(ctx, query, longitude, latitude, maxDistance)
	}
}

func (s *ListingService) GetByDistanceAndSearch(ctx context.Context, latitude, longitude, maxDistance float64, listingType string, searchQuery string) ([]Listing, error) {
	// Sanitize or escape the search query to prevent SQL injection (optional but recommended)
	maxDistance *= 1000
	fmt.Println(longitude, latitude, maxDistance, searchQuery)
	// Check if listingType is valid and apply the relevant filter
	if listingType == "Request" || listingType == "Offer" {
		query := `SELECT * FROM listings 
                 WHERE ST_Distance_Sphere(location, ST_GeomFromText(CONCAT('POINT(', ?, ' ', ?, ')'))) < ? 
                 AND type = ? 
                 AND (title LIKE ? OR description LIKE ?)`
		// Execute the query with the appropriate parameters, including the search query
		return s.queryListings(ctx, query, longitude, latitude, maxDistance, listingType, "%"+searchQuery+"%", "%"+searchQuery+"%")
	} else {
		query := `SELECT * FROM listings 
                 WHERE ST_Distance_Sphere(location, ST_GeomFromText(CONCAT('POINT(', ?, ' ', ?, ')'))) < ? 
                 AND (title LIKE ? OR description LIKE ?)`
		// Execute the query with the search filter
		return s.queryListings(ctx, query, longitude, latitude, maxDistance, "%"+searchQuery+"%", "%"+searchQuery+"%")
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
