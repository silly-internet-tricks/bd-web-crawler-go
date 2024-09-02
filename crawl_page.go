package main

import (
	"fmt"
	"net/url"
	"strings"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	defer cfg.mu.Unlock()
	cfg.mu.Lock()
	_, ok := cfg.pages[normalizedURL]
	if ok {
		cfg.pages[normalizedURL]++
		return
	}

	cfg.pages[normalizedURL] = 1
	isFirst = true
	return
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	cfg.concurrencyControl <- struct{}{}

	cfg.mu.Lock()
	lenPages := len(cfg.pages)
	cfg.mu.Unlock()
	if lenPages >= cfg.maxPages {
		return
	}

	currentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseURL.Host != currentUrl.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("now crawling new page: %v\n", normalizedURL)
	html, err := getHTML("https://" + normalizedURL)
	if err != nil {
		if strings.HasPrefix(err.Error(), "wrong content type") {
			err = nil
		}

		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
		if err != nil {
			return
		}
	}
}
