package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandlerReturnsOK(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	healthHandler(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, response.StatusCode)
	}

	expectedContentType := "application/json"
	actualContentType := response.Header.Get("Content-Type")

	if actualContentType != expectedContentType {
		t.Errorf(
			"expected Content-Type %q, got %q",
			expectedContentType,
			actualContentType,
		)
	}

	var body map[string]string

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}

	if body["status"] != "healthy" {
		t.Errorf("expected status %q, got %q", "healthy", body["status"])
	}
}
