package main

import (
	"diary/internal/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

var srv *httptest.Server
var recordId int64

const (
	testLogin       = "nikita"
	testPassword    = "m5bg8"
	contentTypeJson = "application/json"
)

func TestMain(m *testing.M) {
	serveMux := createServeMux()
	srv = httptest.NewServer(serveMux)
	defer srv.Close()

	err := database.Open()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	m.Run()
}

func TestCreateRecords(t *testing.T) {
	url, _ := url.JoinPath(srv.URL, "records")
	client := srv.Client()

	request, _ := http.NewRequest(http.MethodPost, url, nil)
	t.Run("requires content-type json", func(t *testing.T) {
		response, _ := client.Do(request)
		if response.StatusCode != http.StatusUnsupportedMediaType {
			t.Errorf(
				"status code = %d; want %d",
				response.StatusCode,
				http.StatusUnsupportedMediaType,
			)
		}
	})

	request.Header.Set("Content-Type", contentTypeJson)
	t.Run("requires credentials", func(t *testing.T) {
		response, _ := client.Do(request)
		if response.StatusCode != http.StatusUnauthorized {
			t.Errorf(
				"status code = %d; want %d",
				response.StatusCode,
				http.StatusUnauthorized,
			)
		}
	})

	request.SetBasicAuth("invaliduser", "invalidpassword")
	t.Run("requires right credentials", func(t *testing.T) {
		response, _ := client.Do(request)
		if response.StatusCode != http.StatusUnauthorized {
			t.Errorf(
				"status code = %d; want %d",
				response.StatusCode,
				http.StatusUnauthorized,
			)
		}
	})

	request.SetBasicAuth(testLogin, testPassword)
	badRequestJsons := map[string]string{
		"invalid field spelling": `{"content1": "test record"}`,
		"excess fields":          `{"content": "test record", "extraField": "test"}`,
		"missing field":          `{}`,
		"invalid type - number":  `{"content": 1}`,
		"invalid type - null":    `{"content": null}`,
	}
	for name, json := range badRequestJsons {
		t.Run(name, func(t *testing.T) {
			body := strings.NewReader(json)
			request, _ = http.NewRequest(http.MethodPost, url, body)
			request.SetBasicAuth(testLogin, testPassword)
			request.Header.Set("Content-Type", contentTypeJson)
			response, _ := client.Do(request)
			if response.StatusCode != http.StatusBadRequest {
				t.Errorf("status code = %d; want %d", response.StatusCode, http.StatusBadRequest)
			}
		})
	}

	goodJson := `{"content": "test record"}`
	t.Run("adds new record", func(t *testing.T) {
		body := strings.NewReader(goodJson)
		request, _ = http.NewRequest(http.MethodPost, url, body)
		request.SetBasicAuth(testLogin, testPassword)
		request.Header.Set("Content-Type", contentTypeJson)
		response, _ := client.Do(request)
		if response.StatusCode != http.StatusOK {
			t.Errorf("status code = %d; want %d", response.StatusCode, http.StatusOK)
		}

		contentType := response.Header.Get("Content-Type")
		if contentType != contentTypeJson {
			t.Errorf("invalid content type %s; want %s", contentType, contentTypeJson)
		}

		record := addRecordJsonRes{}
		err := json.NewDecoder(response.Body).Decode(&record)
		if err != nil {
			t.Fatalf("error decoding json: %s", err)
		}

		now := time.Now()
		if now.Unix()-record.CreatedAt.Unix() > 5 {
			t.Errorf("wrong creation time: %v", record.CreatedAt)
		}

		recordId = record.Id
	})
}

func TestGetRecords(t *testing.T) {
	url, _ := url.JoinPath(srv.URL, "records")
	client := srv.Client()
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.SetBasicAuth(testLogin, testPassword)
	response, _ := client.Do(request)
	if response.StatusCode != http.StatusOK {
		t.Errorf("response.StatusCode = %d, want %d", response.StatusCode, http.StatusOK)
	}

	records := []struct {
		Id        int64
		Content   string
		CreatedAt time.Time
	}{}
	err := json.NewDecoder(response.Body).Decode(&records)
	if err != nil {
		t.Errorf("error decoding json: %s", err)
	}

	if len(records) < 1 {
		t.Errorf("len(records) = %d; want > 0", len(records))
	}

	idFound := false
	for _, record := range records {
		if record.Id == recordId {
			idFound = true
			break
		}
	}
	if !idFound {
		t.Errorf(`record with id = %d was not added`, recordId)
	}
}

func TestDeleteRecords(t *testing.T) {
	url, _ := url.JoinPath(srv.URL, "records", strconv.FormatInt(recordId, 10))
	request, _ := http.NewRequest(http.MethodDelete, url, nil)
	request.SetBasicAuth(testLogin, testPassword)

	response, _ := srv.Client().Do(request)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("couldn't delete record with id = %d", recordId)
	}

	response, _ = srv.Client().Do(request)
	if response.StatusCode != http.StatusForbidden {
		t.Fatalf("response.StatusCode = %d; want %d", response.StatusCode, http.StatusForbidden)
	}
}
