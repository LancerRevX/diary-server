package database

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

type User struct {
	Login string
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

func (user *User) HasRecord(id int64) (bool, error) {
	db, err := open()
	if err != nil {
		return false, err
	}
	defer db.Close()

	row := db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM records WHERE id = $1 AND user_login = $2)",
		id,
		user.Login,
	)
	var result bool
	err = row.Scan(&result)
	if err != nil {
		return false, err
	}

	return result, nil
}
