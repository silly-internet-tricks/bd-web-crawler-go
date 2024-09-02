package main

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
	"strings"
)

func getAnchors(nodes *html.Node) (anchors []html.Node) {
	if nodes.Type == html.ElementNode {
		if nodes.DataAtom == atom.A {
			anchors = append(anchors, *nodes)
		}

	}

	if nodes.NextSibling != nil {
		anchors = append(anchors, getAnchors(nodes.NextSibling)...)
	}

	if nodes.FirstChild != nil {
		anchors = append(anchors, getAnchors(nodes.FirstChild)...)
	}

	return
}

func getURLsFromHTML(htmlBody, rawBaseURL string) (urls []string, err error) {
	nodes, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return
	}

	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	anchors := getAnchors(nodes)

	for _, anchor := range anchors {
		for _, attr := range anchor.Attr {
			if attr.Key == "href" {
				var u *url.URL
				u, err = url.Parse(attr.Val)
				if err != nil {
					return
				}

				if u.Host == "" {
					u = baseUrl.JoinPath(u.String())
					urls = append(urls, u.String())
					break
				}

				urls = append(urls, u.String())
				break
			}
		}
	}

	return
}
