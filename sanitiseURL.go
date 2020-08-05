package main

import (
	"fmt"
	url "net/url"
)

// SanitiseURL makes sure that the string supplied is a valid URL
func SanitiseURL(rawURL string) (*url.URL, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return parsedURL, err
	}

	fmt.Println(parsedURL)
	return parsedURL, err
}
