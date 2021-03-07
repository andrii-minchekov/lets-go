package impl

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func connect(dsn string) *sql.DB {

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

type Database struct {
	*sql.DB
}
