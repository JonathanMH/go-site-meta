package main

import (
	"fmt"
	"testing"
)

func TestSanitiseURL(t *testing.T) {
	fmt.Println("TestSanitiseURL")
	validURL, validErr := SanitiseURL("http://jonathanmh.com")

	if validURL.Path != "" {
		t.Error("incorrect path")
	}

	if validErr != nil {
		t.Error("Valid URL flagged as invalid")
	}

	fmt.Println()

	invalidURL, invalidErr := SanitiseURL("justastring")

	fmt.Println(invalidURL)

	if invalidErr == nil {
		t.Error("Invalid URL not flagged as invalid")
	}

	stripFragment, invalidErr := SanitiseURL("http://jonathanmh.com#nothing-to-see")

	fmt.Println("stripFragment", stripFragment)

	if invalidErr == nil {
		t.Error("Invalid URL not flagged as invalid")
	}

}
