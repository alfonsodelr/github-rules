package clients

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	client := NewClient("1.0")
	expected := "1.0"
	if client.GetVersion() != expected {
		t.Errorf("expected version %q, got %q", expected, client.GetVersion())
	}
}

func TestEmpty(t *testing.T) {
}
