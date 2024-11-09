package Services

import (
	"database/sql"
	"time"
)

type Transaction struct {
	TransactionID      int       `json:"transaction_id"`
	UserOfferedID      int       `json:"user_offered_id"`
	UserOfferingID     int       `json:"user_offering_id"`
	ListingID          int       `json:"listing_id"`
	PriceWithCurrency  string    `json:"price_with_currency"`
	DateCreated        time.Time `json:"date_created"`   // Format as "2006-01-02 15:04:05"
	JobStartDate       time.Time `json:"job_start_date"` // Format as "2006-01-02"
	JobEndDate         time.Time `json:"job_end_date"`   // Format as "2006-01-02"
	DetailsFromOffered string    `json:"details_from_offered"`
	DetailsOffering    string    `json:"details_offering"`
}

type TransactionService struct {
	db *sql.DB
}
