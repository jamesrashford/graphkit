package io_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gio "github.com/jamesrashford/graphkit/io"
)

// TestIsURL checks different cases for the IsURL function
func TestIsURL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"ftp://example.com", false}, // Only http and https should be valid
		{"example.com", false},       // Missing scheme
		{"", false},                  // Empty string
		{"random text", false},       // Not a URL
	}

	for _, test := range tests {
		result := gio.IsURL(test.input)
		if result != test.expected {
			t.Errorf("IsURL(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// TestReadUrl uses an HTTP test server to simulate responses
func TestReadUrl(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ok" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, world!"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tests := []struct {
		url         string
		expectError bool
		expectedStr string
	}{
		{server.URL + "/ok", false, "Hello, world!"},
		{server.URL + "/notfound", true, ""},
		{"http://invalid.url", true, ""},
	}

	for _, test := range tests {
		reader, err := gio.ReadUrl(test.url)

		if test.expectError {
			if err == nil {
				t.Errorf("ReadUrl(%q) expected error but got none", test.url)
			}
		} else {
			if err != nil {
				t.Errorf("ReadUrl(%q) unexpected error: %v", test.url, err)
				continue
			}
			// Read from the reader
			body, _ := io.ReadAll(reader)
			if strings.TrimSpace(string(body)) != test.expectedStr {
				t.Errorf("ReadUrl(%q) = %q; want %q", test.url, body, test.expectedStr)
			}
		}
	}
}
