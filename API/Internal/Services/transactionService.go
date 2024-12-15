package Services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Transaction struct {
	TransactionID       int     `json:"transaction_id"`
	UserOfferedID       int     `json:"user_offered_id"`
	UserOfferingID      int     `json:"user_offering_id"`
	ListingID           int     `json:"listing_id"`
	Price               float64 `json:"price_with_currency"`
	CurrencyCode        string  `json:"currency_code"`
	DateCreated         string  `json:"date_created"`   // Format as "2006-01-02 15:04:05"
	JobStartDate        string  `json:"job_start_date"` // Format as "2006-01-02"
	JobEndDate          string  `json:"job_end_date"`   // Format as "2006-01-02"
	DetailsFromOffered  string  `json:"details_from_offered"`
	DetailsFromOffering string  `json:"details_from_offering"`
	Status              string  `json:"status"`
}

type TransactionService struct {
	db *sql.DB
}

// Reusable function to query transactions based on different conditions
func (t *TransactionService) queryTransaction(ctx context.Context, query string, args ...interface{}) ([]Transaction, error) {
	rows, err := t.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve transaction: %w", err)
	}
	defer rows.Close()

	var transactions []Transaction

	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.TransactionID, &transaction.UserOfferedID, &transaction.UserOfferingID,
			&transaction.ListingID, &transaction.Price, &transaction.DateCreated, &transaction.JobStartDate,
			&transaction.JobEndDate, &transaction.DetailsFromOffered, &transaction.DetailsFromOffering,
			&transaction.CurrencyCode, &transaction.Status); err != nil {
			return nil, fmt.Errorf("could not scan transaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over transactions: %v", err)
	}
	if len(transactions) == 0 {
		return []Transaction{}, nil
	}
	return transactions, nil
}

func (t *TransactionService) Create(ctx context.Context, transaction *Transaction) (Transaction, error) {

	query := `INSERT INTO transactions (user_offered_id, user_offering_id, listing_id, 
                          price, job_start_date, job_end_date, details_from_offered, 
                          currency) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := t.db.ExecContext(ctx, query, transaction.UserOfferedID, transaction.UserOfferingID,
		transaction.ListingID, transaction.Price, transaction.JobStartDate, transaction.JobEndDate,
		transaction.DetailsFromOffered, transaction.CurrencyCode)

	if err != nil {
		return Transaction{}, fmt.Errorf("could not create transaction: %w", err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		return Transaction{}, fmt.Errorf("could not get last insert ID: %w", err)
	}

	return Transaction{}, nil
}

func (t *TransactionService) GetByID(ctx context.Context, transactionID int) (Transaction, error) {
	query := `SELECT * FROM transactions WHERE transaction_id = ?`

	transactions, err := t.queryTransaction(ctx, query, transactionID)
	if err != nil {
		return Transaction{}, fmt.Errorf("could not retrieve transaction by ID: %w", err)
	}
	if len(transactions) == 0 {
		return Transaction{}, errors.New("transaction not found")
	}
	return transactions[0], nil
}

func (t *TransactionService) GetByOfferedUserAndStatus(ctx context.Context, offeredUserID int, status string) ([]Transaction, error) {
	// Valid statuses: "Pending", "Accepted", "Completed"
	var query string
	var transactions []Transaction
	var err error
	if status == "Pending" || status == "Accepted" || status == "Completed" {
		query = `SELECT * FROM transactions WHERE user_offered_id = ? AND status = ?`

		transactions, err = t.queryTransaction(ctx, query, offeredUserID, status)
	} else {
		query = `SELECT * FROM transactions WHERE user_offered_id = ?`
		transactions, err = t.queryTransaction(ctx, query, offeredUserID)
	}

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions for offered user ID %d with status %q: %w", offeredUserID, status, err)
	}
	return transactions, nil
}

func (t *TransactionService) GetByOfferingUserAndStatus(ctx context.Context, offeringUserID int, status string) ([]Transaction, error) {
	// Valid statuses: "Pending", "Accepted", "Completed"
	var query string
	var transactions []Transaction
	var err error
	if status == "Pending" || status == "Accepted" || status == "Completed" {
		query = `SELECT * FROM transactions WHERE user_offering_id = ? AND status = ?`

		transactions, err = t.queryTransaction(ctx, query, offeringUserID, status)
	} else {
		query = `SELECT * FROM transactions WHERE user_offering_id = ?`

		transactions, err = t.queryTransaction(ctx, query, offeringUserID)
	}
	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions for offering user ID %d with status %q: %w", offeringUserID, status, err)
	}
	return transactions, nil
}

func (t *TransactionService) GetByListingAndStatus(ctx context.Context, listingID int, status string) ([]Transaction, error) {
	// Valid statuses: "Pending", "Accepted", "Completed"
	var query string
	var transactions []Transaction
	var err error
	if status == "Pending" || status == "Accepted" || status == "Completed" {
		query = `SELECT * FROM transactions WHERE listing_id = ? AND status = ?`

		transactions, err = t.queryTransaction(ctx, query, listingID, status)
	} else {
		query = `SELECT * FROM transactions WHERE listing_id = ?`

		transactions, err = t.queryTransaction(ctx, query, listingID)
	}

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions for listing ID %d: %w", listingID, err)
	}
	return transactions, nil
}

func (t *TransactionService) Delete(ctx context.Context, transactionID int) error {
	query := "DELETE FROM transactions WHERE transaction_id = ?"

	result, err := t.db.ExecContext(ctx, query, transactionID)
	if err != nil {
		return fmt.Errorf("could not delete transaction: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return errors.New("no transaction found to delete")
	}

	return nil
}

func (t *TransactionService) Update(ctx context.Context, id int, transaction Transaction) error {
	query := `UPDATE transactions 
              SET user_offered_id = ?, 
                  user_offering_id = ?, 
                  listing_id = ?, 
                  price = ?, 
                  currency_code = ?, 
                  job_start_date = ?, 
                  job_end_date = ?, 
                  details_from_offered = ?, 
                  details_from_offering = ?, 
                  status = ?
              WHERE transaction_id = ?`

	_, err := t.db.ExecContext(ctx, query,
		transaction.UserOfferedID,
		transaction.UserOfferingID,
		transaction.ListingID,
		transaction.Price,
		transaction.CurrencyCode,
		transaction.JobStartDate,
		transaction.JobEndDate,
		transaction.DetailsFromOffered,
		transaction.DetailsFromOffering,
		transaction.Status,
		id,
	)

	if err != nil {
		return fmt.Errorf("could not update transaction with ID %d: %w", id, err)
	}
	return nil
}
