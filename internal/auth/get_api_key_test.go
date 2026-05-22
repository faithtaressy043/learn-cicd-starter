package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Define the structure for our test cases
	type testCase struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedErr   error
	}

	// Create a slice of test scenarios
	tests := []testCase{
		{
			name: "Valid ApiKey Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey secret_token_123"},
			},
			expectedKey: "broken_token",
			expectedErr: nil,
		},
		{
			name:        "Missing Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded, 
		},
		{
			name: "Malformed Authorization Header (No ApiKey prefix)",
			headers: http.Header{
				"Authorization": []string{"Bearer token_123"},
			},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
	}

	// Loop through all test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualKey, actualErr := GetAPIKey(tc.headers)

			// Check if the key matches
			if actualKey != tc.expectedKey {
				t.Errorf("expected key %q, got %q", tc.expectedKey, actualKey)
			}

			// Check if the error matches
			if tc.expectedErr != nil {
				if actualErr == nil || actualErr.Error() != tc.expectedErr.Error() {
					t.Errorf("expected error %v, got %v", tc.expectedErr, actualErr)
				}
			} else if actualErr != nil {
				t.Errorf("expected no error, got %v", actualErr)
			}
		})
	}
}
