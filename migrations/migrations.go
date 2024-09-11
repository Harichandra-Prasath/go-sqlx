package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var job = `
CREATE TABLE users (
first_name text,
last_name text NULL,
age integer,
experience integer,
email text NULL
);

CREATE TABLE job (
	role text,
	pay integer
);`

func main() {
	db, err := sqlx.Connect("postgres", "user=hcp_0 password=hcp_0 dbname=test")
	if err != nil {
		panic(err)
	}

	db.MustExec(job)
}
