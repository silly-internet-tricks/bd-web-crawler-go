package main

import (
	"net/url"
	"regexp"
)

func normalizeURL(inputURL string) (u string, err error) {
	trailingSlash, err := regexp.Compile(`\/$`)
	if err != nil {
		return
	}

	parsed, err := url.Parse(inputURL)
	u = trailingSlash.ReplaceAllString(parsed.Host+parsed.Path, "")

	return
}
