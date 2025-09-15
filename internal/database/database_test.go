package database_test

import (
	"diary/internal/database"
	"testing"
)

func BenchmarkDatabase(b *testing.B) {
	user := &database.User{Login: "nikita"}
	for b.Loop() {
		record, _ := database.CreateRecord(user, "test record")
		_, _ = database.GetRecords(user)
		_ = database.DeleteRecord(record.Id)
	}
}