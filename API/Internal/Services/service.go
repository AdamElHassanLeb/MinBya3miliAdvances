package Services

import "database/sql"

type Service struct {
	Users interface {
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
