package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlToCheck string) (string, error) {
	var processedURL strings.Builder

	parseURL, err := url.Parse(urlToCheck)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
	}
	// Get Host and Path
	//Make sure host is all lower case.
	processedURL.WriteString(strings.ToLower(parseURL.Hostname()))
	//Remove trailing '/'
	if strings.HasSuffix(parseURL.Path, "/") {
		parseURL.Path = parseURL.Path[0 : len(parseURL.Path)-1]
	}
	processedURL.WriteString(parseURL.Path)

	return processedURL.String(), nil
}
