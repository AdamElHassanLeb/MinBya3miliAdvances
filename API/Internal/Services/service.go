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

func getDB(db *sql.DB) Service {
	return Service{
		Users:       &UserDB{db: db},
		Listings:    &ListingDB{db: db},
		Images:      &ImageDB{db: db},
		Transaction: &TransactionDB{db: db},
	}
}
