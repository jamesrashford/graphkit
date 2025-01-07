package io

import (
	"io"
	"net/http"
	"net/url"
)

func IsUrl(addr string) bool {
	_, err := url.ParseRequestURI(addr)
	if err != nil {
		return false
	}
	return true
}

func ReadUrl(addr string) (io.Reader, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp.Body, nil
}
