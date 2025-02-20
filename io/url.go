package io

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func IsURL(addr string) bool {
	u, err := url.Parse(addr)
	return err == nil && u.Scheme != "" && (u.Scheme == "http" || u.Scheme == "https")
}

func ReadUrl(addr string) (io.Reader, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("received non-OK HTTP status: %d", resp.StatusCode)
	}

	return resp.Body, nil // resp.Body implements io.Reader
}
