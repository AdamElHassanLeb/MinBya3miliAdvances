package Services

import (
	"context"
	"database/sql"
)

type Service struct {
	Users interface {
		GetAll(context.Context) ([]User, error)
		GetById(ctx context.Context, id int) (User, error)
		GetByName(ctx context.Context, name string) ([]User, error)
		Create(context.Context, *User) error
		Update(context.Context, *User) error
		Delete(context.Context, int) (bool, error)
		Auth(context.Context, string, string) (string, User, error)
		GetByPhoneNumber(context.Context, string) (User, error)
	}
	Listings interface {
		Create(context.Context, *Listing) error
		Update(context.Context, *Listing, int) error
		Delete(ctx context.Context, listingID int) error
		GetAll(ctx context.Context, listingType string) ([]Listing, error)
		GetByUserID(ctx context.Context, userID int, listingType string) ([]Listing, error)
		GetByID(ctx context.Context, listingID int) (Listing, error)
		GetBySearch(ctx context.Context, query string, listingType string) ([]Listing, error)
		GetByDistance(ctx context.Context, latitude, longitude, maxDistance float64, listingType string) ([]Listing, error)
		GetByDistanceAndSearch(ctx context.Context, latitude, longitude, maxDistance float64, listingType string, searchQuery string) ([]Listing, error)
		//QueryByLocation(ctx context.Context, latitude, longitude, maxRange float64, listingType string) ([]Listing, error)
		GetByDateCreatedDescending(ctx context.Context, listingType string) ([]Listing, error)
		GetByDateCreatedAndSearchDescending(ctx context.Context, query string, listingType string) ([]Listing, error)
	}
	Images interface {
		AddImage(ctx context.Context, url string, userID int, listingID int) (int, error)
		UpdateImageProfileStatus(ctx context.Context, imageID int, showOnProfile bool) error
		UpdateImageProfilePictureStatus(ctx context.Context, imageID int, user_id int) error
		DeleteImage(context.Context, int) error
		GetImageByID(context.Context, int) (Image, error)
		GetImagesByListingID(context.Context, int) ([]Image, error)
		GetImagesByUserID(context.Context, int) ([]Image, error)
		GetImagesByUserProfile(context.Context, int) ([]Image, error)
	}
	Transaction interface {
	}
}

func ServiceDB(db *sql.DB) Service {
	return Service{
		Users:       &UserService{db: db},
		Listings:    &ListingService{db: db},
		Images:      &ImageService{db: db},
		Transaction: &TransactionService{db: db},
	}
}
