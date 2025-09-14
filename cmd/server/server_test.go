package main

import (
	"diary/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecords(t *testing.T) {
	srv := httptest.NewServer(nil)
	defer srv.Close()

	response, err := srv.Client().Get(srv.URL + "/records")
	if err != nil {
		t.Errorf("couldn't make request: %s", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("response.StatusCode = %d, want %d", response.StatusCode, http.StatusOK)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("couldn't read body: %s", err)
	}
	records := []models.Record{}
	err = json.Unmarshal(body, &records)
	if err != nil {
		t.Errorf("error unmarshalling json: %s", err)
	}

	if len(records) < 1 {
		t.Errorf("len(records) = %d; want > 0", len(records))
	}

	if records[0].Content != "Hello,1 World!" {
		t.Errorf("records[0].Content = \"%s\"; want \"Hello, World!\"", records[0].Content)
	}
}
