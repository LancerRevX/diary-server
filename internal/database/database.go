package database

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func open() (*sql.DB, error) {
	databaseUrl := "postgres://nikita@localhost/diary"
	db, err := sql.Open("pgx", databaseUrl)
	return db, err
}

type User struct {
	Login string
}

type Record struct {
	Id      int
	Content string
}

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidPassword = errors.New("invalid password")

func ValidateCredentials(login string, password string) (*User, error) {
	db, err := open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(
		"SELECT password_hash FROM users WHERE login = $1",
		login,
	)
	var dbPasswordHash string
	err = row.Scan(&dbPasswordHash)
	switch err {
	case nil:
		break
	case sql.ErrNoRows:
		return nil, ErrUserNotFound
	default:
		return nil, err
	}

	passwordHash := md5.Sum([]byte(password))
	passwordHashStr := hex.EncodeToString(passwordHash[:])
	if passwordHashStr != dbPasswordHash {
		return nil, ErrInvalidPassword
	}

	result := User{Login: login}
	return &result, nil
}

func GetRecordsForUser(user *User) ([]Record, error) {
	db, err := open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT id, content FROM records WHERE user_login = $1",
		user.Login,
	)
	if err != nil {
		return nil, err
	}

	result := []Record{}
	for rows.Next() {
		record := Record{}
		err = rows.Scan(&record.Id, &record.Content)
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}

	return result, nil
}
