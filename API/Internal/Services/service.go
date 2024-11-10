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
		Auth(context.Context, string, string) (string, error)
	}
	Listings interface {
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
