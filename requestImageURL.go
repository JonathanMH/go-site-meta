package main

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// FindImageURL returns an og:image or a first image on the page
func FindImageURL(pageURL string) (string, error) {
	doc, err := goquery.NewDocument(pageURL)
	if err != nil {
		fmt.Println(err)
	}
	ogImageURL, ogImageErr := findOGImage(doc)
	fmt.Println(ogImageURL, ogImageErr)
	return ogImageURL, nil
}

func findOGImage(pageDocument *goquery.Document) (imageURL string, err error) {
	var ogImageURL string

	pageDocument.Find("meta").Each(func(index int, item *goquery.Selection) {
		if item.AttrOr("property", "") == "og:image" {
			ogImageURL = item.AttrOr("content", "")
			fmt.Println("found og:image: ", ogImageURL)
		}
	})

	if ogImageURL != "" {
		return ogImageURL, nil
	}

	return "", errors.New("could not find an og:image")
}

func findTwitterImage(pageDocument *goquery.Document) (string, error) {
	var twitterImageURL string

	pageDocument.Find("meta").Each(func(index int, item *goquery.Selection) {
		if item.AttrOr("property", "") == "twitter:image" {
			twitterImageURL = item.AttrOr("content", "")
			fmt.Println("found twitter:image: ", twitterImageURL)
		}
	})

	if twitterImageURL != "" {
		return twitterImageURL, nil
	}

	return "", errors.New("could not find an og:image")
}
