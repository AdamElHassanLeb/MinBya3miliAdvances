package Services

import (
	"context"
	"database/sql"
	geo "github.com/paulmach/go.geo"
)

type Service struct {
	Users interface {
		GetAll(context.Context) ([]User, error)
		GetById(ctx context.Context, id int) (User, error)
		GetByName(ctx context.Context, name string) ([]User, error)
		Create(context.Context, *User) error
		Update(context.Context, *User) error
		Delete(context.Context, int) (bool, error)
		Auth(context.Context, string, string) (string, error)
		GetByPhoneNumber(context.Context, string) (User, error)
	}
	Listings interface {
		Create(userID int, title, description string, location *geo.Point, listingType string) (Listing, error)
		Update(listingID int, title, description string, location *geo.Point, listingType string) (Listing, error)
		Delete(listingID int) error
		GetAll(listingType string) ([]Listing, error)
		GetByUserID(userID int, listingType string) ([]Listing, error)
		GetByID(listingID int) (Listing, error)
		GetBySearch(query string, listingType string) ([]Listing, error)
		GetByDistance(latitude, longitude, maxDistance float64, listingType string) ([]Listing, error)
		QueryByLocation(latitude, longitude, maxRange float64, listingType string) ([]Listing, error)
		GetByDateCreatedDescending(listingType string) ([]Listing, error)
		GetByDateCreatedAndSearchDescending(query string, listingType string) ([]Listing, error)
	}
	Images interface {
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
