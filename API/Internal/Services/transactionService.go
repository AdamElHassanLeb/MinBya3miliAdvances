package Services

import "database/sql"

type TransactionDB struct {
	db *sql.DB
}
