package main

import (
	url "net/url"
)

// SanitiseURL makes sure that the string supplied is a valid URL
func SanitiseURL(rawURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return nil, err
	}

	return parsedURL, nil
}
