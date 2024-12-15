package Services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"os"
	"path/filepath"
)

// Image represents an image in the system
// @Description An image associated with a user or listing, with information about its URL, visibility, and creation date.
// @Model
type Image struct {
	// ImageID is the unique identifier for the image
	// @example 123
	ImageID int `json:"image_id"`

	// URL is the link to the image
	// @example "https://example.com/images/123.jpg"
	URL string `json:"url"`

	// UserID is the ID of the user who uploaded the image
	// @example 1
	UserID int `json:"user_id"`

	// ListingID is the ID of the listing to which the image is associated, if applicable
	// Nullable field, can be omitted or null if not associated with a listing
	// @example 101
	ListingID int `json:"listing_id,omitempty"`

	// ShowOnProfile specifies if the image should be displayed on the user's profile
	// @example true
	ShowOnProfile bool `json:"show_on_profile"`

	// DateCreated is the date when the image was uploaded
	// Format: "2006-01-02 15:04:05"
	// This field is omitted in JSON responses
	// @example "2024-12-16 14:30:00"
	DateCreated string `json:"-"` // Exclude from JSON output
}

var imagesDIR string = Env.GetString("SRV_DIR", "") + "/ServerImages"

type ImageService struct {
	db *sql.DB
}

func (s *ImageService) AddImage(ctx context.Context, url string, userID int, listingID int) (int, error) {
	var query string
	var imageID int

	if listingID != 0 {
		// Insert with ListingID if it's provided
		query = `
			INSERT INTO images (url, user_id, listing_id, show_on_profile)
			VALUES (?, ?, ?, ?)
		`
		result, err := s.db.ExecContext(ctx, query, url, userID, listingID, true) // Assuming show_on_profile is true by default
		if err != nil {
			return 0, fmt.Errorf("could not insert image: %v", err)
		}

		// Get the last inserted ID
		imageID64, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("could not get last insert ID: %v", err)
		}
		imageID = int(imageID64)
	} else {
		// Insert without ListingID if it's 0
		query = `
			INSERT INTO images (url, user_id, show_on_profile)
			VALUES (?, ?, ?)
		`
		result, err := s.db.ExecContext(ctx, query, url, userID, true)
		if err != nil {
			return 0, fmt.Errorf("could not insert image: %v", err)
		}

		// Get the last inserted ID
		imageID64, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("could not get last insert ID: %v", err)
		}
		imageID = int(imageID64)
	}

	return imageID, nil
}

func (s *ImageService) GetImageByID(ctx context.Context, imageID int) (Image, error) {
	query := `SELECT image_id, url, user_id, listing_id, show_on_profile, date_created 
	          FROM images WHERE image_id = ?`

	var image Image
	err := s.db.QueryRowContext(ctx, query, imageID).Scan(&image.ImageID, &image.URL, &image.UserID, &image.ListingID, &image.ShowOnProfile, &image.DateCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			return Image{}, fmt.Errorf("image not found")
		}
		return Image{}, fmt.Errorf("could not get image: %v", err)
	}
	return image, nil
}

func (s *ImageService) GetImagesByListingID(ctx context.Context, listingID int) ([]Image, error) {
	query := `SELECT image_id, url, user_id, listing_id, show_on_profile, date_created 
	          FROM images WHERE listing_id = ?`

	rows, err := s.db.QueryContext(ctx, query, listingID)
	if err != nil {
		return nil, fmt.Errorf("could not get images by listing ID: %v", err)
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var image Image
		if err := rows.Scan(&image.ImageID, &image.URL, &image.UserID, &image.ListingID, &image.ShowOnProfile, &image.DateCreated); err != nil {
			return nil, fmt.Errorf("could not scan image: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return images, nil
}

func (s *ImageService) GetImagesByUserID(ctx context.Context, userID int) ([]Image, error) {
	query := `SELECT image_id, url, user_id, listing_id, show_on_profile, date_created 
	          FROM images WHERE user_id = ?`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get images by user ID: %v", err)
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var image Image
		if err := rows.Scan(&image.ImageID, &image.URL, &image.UserID, &image.ListingID, &image.ShowOnProfile, &image.DateCreated); err != nil {
			return nil, fmt.Errorf("could not scan image: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return images, nil
}

func (s *ImageService) GetImagesByUserProfile(ctx context.Context, userID int) ([]Image, error) {
	query := `SELECT image_id, url, user_id, listing_id, show_on_profile, date_created 
	          FROM images WHERE user_id = ? AND show_on_profile = 1`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get images by user profile: %v", err)
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var image Image
		if err := rows.Scan(&image.ImageID, &image.URL, &image.UserID, &image.ListingID, &image.ShowOnProfile, &image.DateCreated); err != nil {
			return nil, fmt.Errorf("could not scan image: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return images, nil
}

func (s *ImageService) UpdateImageProfilePictureStatus(ctx context.Context, imageID int, user_id int) error {
	query := `
		UPDATE users 
		SET profile_image = ? 
		WHERE user_id = ?;`

	_, err := s.db.ExecContext(ctx, query, imageID, user_id)
	if err != nil {
		return fmt.Errorf("could not update image: %v", err)
	}
	return nil
}

func (s *ImageService) UpdateImageProfileStatus(ctx context.Context, imageID int, showOnProfile bool) error {
	query := `
		UPDATE images 
		SET show_on_profile = ? 
		WHERE image_id = ?`

	_, err := s.db.ExecContext(ctx, query, showOnProfile, imageID)
	if err != nil {
		return fmt.Errorf("could not update image: %v", err)
	}
	return nil
}

func (s *ImageService) DeleteImage(ctx context.Context, imageID int) error {
	// Retrieve the URL of the image to delete from the file system
	var url string
	query := `SELECT url FROM images WHERE image_id = ?`
	err := s.db.QueryRowContext(ctx, query, imageID).Scan(&url)
	if err != nil {
		return fmt.Errorf("could not find image: %v", err)
	}

	// Optionally, delete the image file from the server (assuming file path is the full URL)
	err = os.Remove(filepath.Join(imagesDIR, url))
	if err != nil {
		return fmt.Errorf("could not delete image file: %v", err)
	}

	// Now delete the image record from the database
	query = `DELETE FROM images WHERE image_id = ?`
	_, err = s.db.ExecContext(ctx, query, imageID)
	if err != nil {
		return fmt.Errorf("could not delete image record: %v", err)
	}
	return nil
}
