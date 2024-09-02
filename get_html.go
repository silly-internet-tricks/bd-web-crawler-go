package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (html string, err error) {
	sponse, err := http.Get(rawURL)
	if err != nil {
		return
	}

	if sponse.StatusCode >= 400 {
		err = fmt.Errorf("got an error code: %v", sponse.StatusCode)
		return
	}

	contentType := sponse.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		err = errors.New("wrong content type; text/html is required")
		return
	}

	body, err := io.ReadAll(sponse.Body)
	if err != nil {
		return
	}

	html = string(body)

	return
}
