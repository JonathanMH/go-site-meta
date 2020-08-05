package main

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFindImageURL(t *testing.T) {
	fmt.Println("TestFindImageURL")
	myMockPage := `<!DOCTYPE html><html lang=en-US class=no-js><head><meta roperty=og:site_name content=JonathanMH> <meta property=og:image content=https://jonathanmh.com/wp-content/uploads/2020/02/docker-compose-php5-php7-parallel.png> </head></html>`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://jonathanmh.com/not-a-real-url/",
		httpmock.NewStringResponder(200, myMockPage))

	imageURL, _ := FindImageURL("https://jonathanmh.com/not-a-real-url/")

	if imageURL != "https://jonathanmh.com/wp-content/uploads/2020/02/docker-compose-php5-php7-parallel.png" {
		t.Error("imageURL was incorrect")
	}

	fmt.Println(imageURL)
}
