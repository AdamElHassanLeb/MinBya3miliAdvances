package Database

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func DBConnection(addr string) (*sql.DB, error) {

	db, err := sql.Open("mysql", addr)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	log.Print("DB Connected \n")
	return db, nil
}
