package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DB_INIT() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=hcp_0 password=hcp_0 dbname=test")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
