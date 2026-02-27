package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	newMux().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	expected := "Handler default is processing the version request with version 1.0"
	if rr.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rr.Body.String())
	}
}
