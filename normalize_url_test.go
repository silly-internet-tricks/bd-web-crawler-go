package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
		err      error
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "remove trailing slash",
			inputURL: "blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "remove trailing slash and insecure scheme",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "remove insecure scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "handles empty string as url",
			inputURL: "",
			expected: "",
			err:      nil,
		},
		{
			name:     "handles x's as url",
			inputURL: "xxxxXXXxxxxXxxxxXxxxXxxxx",
			expected: "xxxxXXXxxxxXxxxxXxxxXxxxx",
			err:      nil,
		},
		{
			name:     "remove query",
			inputURL: "http://blog.boot.dev/path?search=wagslane",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "remove fragment",
			inputURL: "http://blog.boot.dev/path#main-content",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
		{
			name:     "remove query and fragment",
			inputURL: "http://blog.boot.dev/path?hello=world#second-heading",
			expected: "blog.boot.dev/path",
			err:      nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)

			if tc.err != nil && err == nil {
				t.Errorf("Test %v - '%s' FAIL: missing error: %v", i, tc.name, tc.err)
				return
			}

			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
