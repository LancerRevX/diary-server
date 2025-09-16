package database

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

type User struct {
	Id    int64
	Login string
}

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidPassword = errors.New("invalid password")

func ValidateCredentials(login string, password string) (*User, error) {
	user := User{Login: login}

	row := db.QueryRow(
		"SELECT id, password_hash FROM users WHERE login = $1",
		login,
	)
	var dbPasswordHash string
	err := row.Scan(&user.Id, &dbPasswordHash)
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

	return &user, nil
}
