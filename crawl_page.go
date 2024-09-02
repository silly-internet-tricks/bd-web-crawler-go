package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (p map[string]int, err error) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if baseUrl.Host != currentUrl.Host {
		p = pages
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL]++
		p = pages
		return
	}

	pages[normalizedURL] = 1
	p = pages
	fmt.Printf("now crawling new page: %v\n", normalizedURL)
	html, err := getHTML("https://" + normalizedURL)
	if err != nil {
		if strings.HasPrefix(err.Error(), "wrong content type") {
			err = nil
		}

		return
	}

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		return
	}

	for _, url := range urls {
		p, err = crawlPage(rawBaseURL, url, p)
		if err != nil {
			return
		}
	}

	return
}
