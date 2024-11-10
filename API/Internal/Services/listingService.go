package Services

import (
	"database/sql"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Utils"
	geo "github.com/paulmach/go.geo"
	"time"
)

// Listing represents the listing structure
type Listing struct {
	ListingID   int        `json:"listing_id"`
	Type        string     `json:"type"`     // Enum: 'Request' or 'Offer'
	Location    *geo.Point `json:"location"` // Custom handling may be needed for POINT
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DateCreated time.Time  `json:"date_created"` // Format as "2006-01-02 15:04:05"
	Active      bool       `json:"active"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
}

// ListingService is the service layer for listing-related operations
type ListingService struct {
	db *sql.DB
}

// ValidateCoordinates will check the longitude and latitude by calling the ReverseGeocode function
// from the Utils package to validate the coordinates and get the city and country.
func (s *ListingService) ValidateCoordinates(lat, lon float64) (string, string, error) {
	// Call the ReverseGeocode function from the Utils package to get city and country
	city, country, err := Utils.ReverseGeocode(lat, lon)
	if err != nil {
		return "", "", fmt.Errorf("error validating coordinates: %w", err)
	}
	return city, country, nil
}

// Reusable function to query listings based on different conditions
func (s *ListingService) queryListings(query string, args ...interface{}) ([]Listing, error) {
	rows, err := s.db.Query(query, args...)
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

	return listings, nil
}

// Create a new listing with city and country based on coordinates
func (s *ListingService) Create(userID int, title, description string, location *geo.Point, listingType string) (Listing, error) {
	// Validate the coordinates (longitude, latitude) using the ReverseGeocode function
	city, country, err := s.ValidateCoordinates(location.Lat(), location.Lng())
	if err != nil {
		return Listing{}, err
	}

	// Prepare the insert query
	query := `
		INSERT INTO listings (type, location, user_id, title, description, city, country)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		RETURNING listing_id, type, location, user_id, title, description, date_created, active, city, country
	`

	// Execute the query
	row := s.db.QueryRow(query, listingType, location, userID, title, description, city, country)

	// Scan the result into a Listing struct
	var listing Listing
	if err := row.Scan(&listing.ListingID, &listing.Type, &listing.Location, &listing.UserID, &listing.Title, &listing.Description, &listing.DateCreated, &listing.Active, &listing.City, &listing.Country); err != nil {
		return Listing{}, fmt.Errorf("could not create listing: %v", err)
	}

	return listing, nil
}

// Update a listing
func (s *ListingService) Update(listingID int, title, description string, location *geo.Point, listingType string) (Listing, error) {
	// Validate the coordinates (longitude, latitude) using the ReverseGeocode function
	city, country, err := s.ValidateCoordinates(location.Lat(), location.Lng())
	if err != nil {
		return Listing{}, err
	}

	// Prepare the update query
	query := `
		UPDATE listings
		SET title = ?, description = ?, location = ?, type = ?, city = ?, country = ?
		WHERE listing_id = ?
		RETURNING listing_id, type, location, user_id, title, description, date_created, active, city, country
	`

	// Execute the query
	row := s.db.QueryRow(query, title, description, location, listingType, city, country, listingID)

	// Scan the result into a Listing struct
	var listing Listing
	if err := row.Scan(&listing.ListingID, &listing.Type, &listing.Location, &listing.UserID, &listing.Title, &listing.Description, &listing.DateCreated, &listing.Active, &listing.City, &listing.Country); err != nil {
		return Listing{}, fmt.Errorf("could not update listing: %v", err)
	}

	return listing, nil
}

// Delete a listing
func (s *ListingService) Delete(listingID int) error {
	// Prepare the delete query
	query := `
		DELETE FROM listings WHERE listing_id = ?
	`

	// Execute the query
	result, err := s.db.Exec(query, listingID)
	if err != nil {
		return fmt.Errorf("could not delete listing: %v", err)
	}

	// Check if any row was affected (ensuring the listing exists)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check affected rows: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("listing not found")
	}

	return nil
}

// Get all listings, with the option to filter by type (request, offer, or both)
func (s *ListingService) GetAll(listingType string) ([]Listing, error) {
	var query string
	switch listingType {
	case "Request":
		query = "SELECT * FROM listings WHERE type = 'Request'"
	case "Offer":
		query = "SELECT * FROM listings WHERE type = 'Offer'"
	default:
		query = "SELECT * FROM listings"
	}

	// Use the reusable query function to fetch results
	return s.queryListings(query)
}

// Get listing by user ID, with type filter
func (s *ListingService) GetByUserID(userID int, listingType string) ([]Listing, error) {
	var query string
	switch listingType {
	case "Request":
		query = "SELECT * FROM listings WHERE user_id = ? AND type = 'Request'"
	case "Offer":
		query = "SELECT * FROM listings WHERE user_id = ? AND type = 'Offer'"
	default:
		query = "SELECT * FROM listings WHERE user_id = ?"
	}

	// Use the reusable query function to fetch results
	return s.queryListings(query, userID)
}

// Get listing by ID
func (s *ListingService) GetByID(listingID int) (Listing, error) {
	query := "SELECT * FROM listings WHERE listing_id = ?"
	listings, err := s.queryListings(query, listingID)
	if err != nil {
		return Listing{}, err
	}

	if len(listings) == 0 {
		return Listing{}, fmt.Errorf("listing not found")
	}

	return listings[0], nil
}

// Get listings by title or description with an optional type filter
func (s *ListingService) GetBySearch(query string, listingType string) ([]Listing, error) {
	var queryStr string
	switch listingType {
	case "Request":
		queryStr = "SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?) AND type = 'Request'"
	case "Offer":
		queryStr = "SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?) AND type = 'Offer'"
	default:
		queryStr = "SELECT * FROM listings WHERE (title LIKE ? OR description LIKE ?)"
	}

	// Use the reusable query function to fetch results
	return s.queryListings(queryStr, "%"+query+"%", "%"+query+"%")
}

// Get listings within a given distance from a location
func (s *ListingService) GetByDistance(latitude, longitude, maxDistance float64, listingType string) ([]Listing, error) {
	var query string
	switch listingType {
	case "Request":
		query = `
			SELECT * FROM listings
			WHERE ST_Distance(location, ST_GeomFromText('POINT(? ?)', 4326)) <= ? AND type = 'Request'
		`
	case "Offer":
		query = `
			SELECT * FROM listings
			WHERE ST_Distance(location, ST_GeomFromText('POINT(? ?)', 4326)) <= ? AND type = 'Offer'
		`
	default:
		query = `
			SELECT * FROM listings
			WHERE ST_Distance(location, ST_GeomFromText('POINT(? ?)', 4326)) <= ?
		`
	}

	// Use the reusable query function to fetch results
	return s.queryListings(query, longitude, latitude, maxDistance)
}

// Query listings based on location and a max range for distance, with optional type filter
func (s *ListingService) QueryByLocation(latitude, longitude, maxRange float64, listingType string) ([]Listing, error) {
	var query string
	switch listingType {
	case "Request":
		query = `
			SELECT * FROM listings
			WHERE ST_Distance(location, ST_GeomFromText('POINT(? ?)', 4326)) <= ? AND type = 'Request'
		`
	case "Offer":
		query = `
			SELECT * FROM listings
			WHERE ST_Distance(location, ST_GeomFromText('POINT(? ?)', 4326)) <= ? AND type = 'Offer'
		`
	default:
		query = `
			SELECT * FROM listings
			WHERE ST_Distance(location, ST_GeomFromText('POINT(? ?)', 4326)) <= ?
		`
	}

	// Use the reusable query function to fetch results
	return s.queryListings(query, longitude, latitude, maxRange)
}

// Get listings ordered by date created in descending order, with optional type filter
func (s *ListingService) GetByDateCreatedDescending(listingType string) ([]Listing, error) {
	var query string
	switch listingType {
	case "Request":
		query = "SELECT * FROM listings WHERE type = 'Request' ORDER BY date_created DESC"
	case "Offer":
		query = "SELECT * FROM listings WHERE type = 'Offer' ORDER BY date_created DESC"
	default:
		query = "SELECT * FROM listings ORDER BY date_created DESC"
	}

	// Use the reusable query function to fetch results
	return s.queryListings(query)
}

// Get listings ordered by date created in descending order with search query, and optional type filter
func (s *ListingService) GetByDateCreatedAndSearchDescending(query string, listingType string) ([]Listing, error) {
	var queryStr string
	switch listingType {
	case "Request":
		queryStr = `
			SELECT * FROM listings
			WHERE (title LIKE ? OR description LIKE ?) AND type = 'Request'
			ORDER BY date_created DESC
		`
	case "Offer":
		queryStr = `
			SELECT * FROM listings
			WHERE (title LIKE ? OR description LIKE ?) AND type = 'Offer'
			ORDER BY date_created DESC
		`
	default:
		queryStr = `
			SELECT * FROM listings
			WHERE (title LIKE ? OR description LIKE ?)
			ORDER BY date_created DESC
		`
	}

	// Use the reusable query function to fetch results
	return s.queryListings(queryStr, "%"+query+"%", "%"+query+"%")
}
