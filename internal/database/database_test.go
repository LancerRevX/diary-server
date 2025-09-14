package database

import (
	"testing"
)

func TestAuthenticateUser(t *testing.T) {
	var user *User
	var err error

	user, err = ValidateCredentials("invaliduser", "m5bg8")
	if err != ErrUserNotFound {
		t.Errorf("err = %s; want %s", err, ErrUserNotFound)
	}
	if user != nil {
		t.Errorf("user = %v; want nil", user)
	}

	user, err = ValidateCredentials("nikita", "invalidpassword")
	if err != ErrInvalidPassword {
		t.Errorf("err = %s; want %s", err, ErrInvalidPassword)
	}
	if user != nil {
		t.Errorf("user = %v; want nil", user)
	}

	user, err = ValidateCredentials("nikita", "m5bg8")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	if user == nil {
		t.Fatalf("user = nil; want user")
	}
	if user.Login != "nikita" {
		t.Errorf("user.login = %s; want nikita", user.Login)
	}
}

func TestGetRecords(t *testing.T) {
	user, _ := ValidateCredentials("nikita", "m5bg8")
	records, err := GetRecordsForUser(user)
	if err != nil {
		t.Fatalf("records = nil")
	}
	if len(records) < 1 {
		t.Errorf("len(records) = %d; want > 1", len(records))
	}
}
