package Services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"os"
	"path/filepath"
)

type Image struct {
	ImageID       int    `json:"image_id"`
	URL           string `json:"-"`
	UserID        int    `json:"user_id"`
	ListingID     int    `json:"listing_id,omitempty"` // Nullable
	ShowOnProfile bool   `json:"show_on_profile"`
	DateCreated   string `json:"-"`
}

var imagesDIR string = Env.GetString("SRV_DIR", "") + "/ServerImages"

type ImageService struct {
	db *sql.DB
}

func (s *ImageService) AddImage(ctx context.Context, url string, userID int, listingID int) error {
	var query string
	var err error

	if listingID != 0 {
		// Insert with ListingID if it's provided
		query = `
			INSERT INTO images (url, user_id, listing_id, show_on_profile)
			VALUES (?, ?, ?, ?)
		`
		_, err = s.db.ExecContext(ctx, query, url, userID, listingID, true) // Assuming show_on_profile is true by default
	} else {
		// Insert without ListingID if it's 0
		query = `
			INSERT INTO images (url, user_id, show_on_profile)
			VALUES (?, ?, ?)
		`
		_, err = s.db.ExecContext(ctx, query, url, userID, true) // No ListingID field
	}

	if err != nil {
		return fmt.Errorf("could not insert image: %v", err)
	}

	return nil
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

func (s *ImageService) UpdateImage(ctx context.Context, imageID int, showOnProfile bool) error {
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
