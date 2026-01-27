package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey second")

	got, err := GetAPIKey(h)
	want := "second"

	if got != want || err != nil {
		t.Errorf("Test failed.  Error: %v", err)
	}

}
