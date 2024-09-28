package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey someApiKeyHereYayyyy")
	have, _ := GetAPIKey(header)
	want := "someApiKeyHereYayyyy"
	if have != want {
		t.Fatalf("expected %v, got %v", want, have)
	}
}
