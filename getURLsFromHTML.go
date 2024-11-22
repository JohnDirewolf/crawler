package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func processNode(node *html.Node, URLstring []string) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				URLstring = append(URLstring, attribute.Val)
				break
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		URLstring = processNode(child, URLstring)
	}
	return URLstring
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var URLs []string

	htmlReader := strings.NewReader(htmlBody)
	htmlTree, err := html.Parse(htmlReader)
	if err != nil {
		fmt.Printf("Error Parsing HTML: %v", err)
	}
	URLs = processNode(htmlTree, URLs)
	//Run through URLs found and if there are relative URLs, add the rawBaseURL.
	if len(URLs) == 0 {
		return []string{}, nil
	}
	for i, URL := range URLs {
		if strings.HasPrefix(URL, "/") {
			URLs[i] = rawBaseURL + URL
		}
	}
	return URLs, nil
}
