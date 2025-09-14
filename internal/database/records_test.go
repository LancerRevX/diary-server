package database

import "testing"

var user *User = &User{Login: "nikita"}
var createdRecordId int64

func TestRecords(t *testing.T) {
	t.Run("create records", testCreateRecords)
	t.Run("get records", testGetRecords)
	t.Run("delete records", testDeleteRecord)
}

func testCreateRecords(t *testing.T) {
	content := "Hello, World!"
	record, err := CreateRecord(user, content)
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

func testGetRecords(t *testing.T) {
	records, err := GetRecords(user)
	if err != nil {
		t.Fatal(err)
	}
	if records == nil {
		t.Fatalf("records = nil")
	}
	if len(records) < 1 {
		t.Errorf("len(records) = %d; want > 1", len(records))
	}
}

func testDeleteRecord(t *testing.T) {
	err := DeleteRecord(createdRecordId)
	if err != nil {
		t.Fatal(err)
	}

	records, _ := GetRecords(user)
	for _, record := range records {
		if record.Id == createdRecordId {
			t.Errorf("record with id = %d not deleted", createdRecordId)
		}
	}
}
