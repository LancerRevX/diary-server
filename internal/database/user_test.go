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

func TestHasRecord(t *testing.T) {
	user := &User{Login: "nikita"}
	record, _ := CreateRecord(user, "test record")
	t.Run("test has record", func(t *testing.T) {
		result, err := user.HasRecord(record.Id)
		if err != nil {
			t.Fatal(err)
		}
		if !result {
			t.Errorf("%v.HasRecord(%v) = false", user, record)
		}

		anotherUser := &User{Login: "test-user"}
		result, err = anotherUser.HasRecord(record.Id)
		if err != nil {
			t.Fatal(err)
		}
		if result {
			t.Errorf("%v.HasRecord(%v) = true", anotherUser, record)
		}
	})
	DeleteRecord(record.Id)
}
