package db

import (
	"context"
	"database/sql"
	"github.com/andrii-minchekov/lets-go/app/impl/cfg"
	"github.com/andrii-minchekov/lets-go/app/impl/db/models"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"log"
	"sync"
)

var once sync.Once

var db *sql.DB

func connect(dsn string) *sql.DB {
	once.Do(func() {
		dbo, err := sql.Open("postgres", dsn)
		db = dbo
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}
	})
	return db
}

type Database struct {
	*sql.DB
}

func GetDatabase() Database {
	return Database{connect(cfg.FlagConfig.DSN())}
}

func CleanupDb() {
	boil.DebugMode = true
	database := GetDatabase()
	models.Snippets().DeleteAll(context.Background(), database)
	models.Users().DeleteAll(context.Background(), database)
	queries.RawG("ALTER SEQUENCE snippets_id_seq RESTART WITH 1;").Exec(database)
	queries.RawG("ALTER SEQUENCE users_id_seq RESTART WITH 1;").Exec(database)
}
