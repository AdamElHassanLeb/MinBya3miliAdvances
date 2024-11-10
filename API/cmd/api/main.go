package main

import (
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Database"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"log"
)

func main() {

	config := config{
		address: Env.GetString("ADDR", ":"),
		db: dbConfig{
			addr: Env.GetString("DB_ADDR", "")},
	}

	db, err := Database.DBConnection(config.db.addr)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	Service := Services.ServiceDB(db)

	app := &application{
		config:  config,
		Service: Service,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))

}
