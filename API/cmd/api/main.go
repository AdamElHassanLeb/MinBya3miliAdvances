package main

import (
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/APP/Internal/Database"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/APP/Internal/env"
	"log"
)

func main() {

	config := config{
		address: env.GetString("ADDR", ":"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "")},
	}

	db, err := Database.DBConnection(config.db.addr)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	//store := Services.NewStorage(db)

	app := &application{
		config: config,
		//store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
