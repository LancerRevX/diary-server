package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func open() (*sql.DB, error) {
	databaseUrl := "postgres://nikita@localhost/diary"
	db, err := sql.Open("pgx", databaseUrl)
	return db, err
}


