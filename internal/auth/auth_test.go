package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	correctHeader := make(http.Header)
	correctHeader.Add("authorization", "ApiKey apiKey")

	noAuthHeader := make(http.Header)

	missingApiKey := make(http.Header)
	missingApiKey.Add("authorization", "ApiKey")

	incorrectAuthAttr := make(http.Header)
	incorrectAuthAttr.Add("authorization", "IncorrectAttr apiKey")

	tests := map[string]struct {
		header     http.Header
		wantApiKey string
		wantErr    bool
	}{
		"Correct Auth Header": {
			header:     correctHeader,
			wantApiKey: "apiKey",
			wantErr:    false,
		},
		"No Auth Header": {
			header:     noAuthHeader,
			wantApiKey: "",
			wantErr:    true,
		},
		"Missing apiKey": {
			header:     missingApiKey,
			wantApiKey: "",
			wantErr:    true,
		},
		"Incorrect Auth attr": {
			header:     incorrectAuthAttr,
			wantApiKey: "",
			wantErr:    true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotApiKey, gotErr := GetAPIKey(tc.header)
			if (gotErr != nil) != tc.wantErr {
				t.Errorf("expected error to be: %v, got: %v", tc.wantErr, gotErr)
			}
			if gotApiKey != tc.wantApiKey {
				t.Errorf("expected apiKey to be: '%s', got: %s", tc.wantApiKey, gotApiKey)
			}
		})
	}
}
