package database

import "testing"

const userId = 1

var createdRecordId int64

func TestMain(m *testing.M) {
	err := Open()
	if err != nil {
		panic(err)
	}
	defer Close()

	m.Run()
}

func TestCreateRecords(t *testing.T) {
	content := "Hello, World!"
	recordId, err := CreateRecord(userId, content, nil)
	if err != nil {
		t.Fatal(err)
	}

	record, err := GetRecordById(recordId)
	if err != nil {
		t.Fatal(err)
	}
	if record == nil {
		t.Fatalf("record = nil")
	}
	if record.Content != content {
		t.Errorf("record.Content = %s; want %s", record.Content, content)
	}

	createdRecordId = record.Id
}

func TestGetRecords(t *testing.T) {
	records, err := GetRecordsByUser(userId)
	if err != nil {
		t.Fatal(err)
	}
	if records == nil {
		t.Fatalf("records = nil")
	}
	if len(records) < 1 {
		t.Errorf("len(records) = %d; want > 1", len(records))
	}

	for _, record := range records {
		if record.Content == "" {
			t.Errorf(`record.Content == ""`)
		}
	}
}

func TestUserHasRecord(t *testing.T) {
	// user := &User{Login: "nikita"}
	// recordId, _ := CreateRecord(user.Id, "test record", nil)
	// t.Run("test has record", func(t *testing.T) {
	// 	result, err := user.HasRecord(record.Id)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	if !result {
	// 		t.Errorf("%v.HasRecord(%v) = false", user, record)
	// 	}

	// 	anotherUser := &User{Login: "test-user"}
	// 	result, err = anotherUser.HasRecord(record.Id)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	if result {
	// 		t.Errorf("%v.HasRecord(%v) = true", anotherUser, record)
	// 	}
	// })
	// DeleteRecord(record.Id)
}

func TestDeleteRecord(t *testing.T) {
	err := DeleteRecord(createdRecordId)
	if err != nil {
		t.Fatal(err)
	}

	records, _ := GetRecordsByUser(userId)
	for _, record := range records {
		if record.Id == createdRecordId {
			t.Errorf("record with id = %d not deleted", createdRecordId)
		}
	}
}
