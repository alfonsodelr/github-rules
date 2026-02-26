package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleVersionRequest(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler directly
	handler := http.HandlerFunc(HandlerVersionRequest)
	handler.ServeHTTP(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	// Check response body
	expected := "Handler default is processing the version request with version 1.0"
	if rr.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}
