package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	got, err := GetAPIKey(http.Header{"Authorization": []string{"ApiKey testKey"}})
	if err != nil {
		t.Fatalf("expected: no error, got: %v", err)
	}
	if got != "testKey" {
		t.Fatalf("expected: %v, got: %v", "testKey", got)
	}
}

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_MalformedAuthHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{"Authorization": []string{"Bearer testKey"}})
	if err == nil || err.Error() != "malformed authorization header" {
		t.Fatalf("expected: %v, got: %v", "malformed authorization header", err)
	}
}
