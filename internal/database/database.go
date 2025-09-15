package database

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func Open() error {
	databaseUrl := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("pgx", databaseUrl)
	return err
}

func Close() {

}
