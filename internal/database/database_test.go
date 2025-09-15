package database_test

import (
	"diary/internal/database"
	"testing"
)

func BenchmarkDatabase(b *testing.B) {
	err := database.Open()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	const userId = 1

	for b.Loop() {
		recordId, _ := database.CreateRecord(userId, "test record", nil)
		_, _ = database.GetRecords(userId)
		_ = database.DeleteRecord(recordId)
	}
}