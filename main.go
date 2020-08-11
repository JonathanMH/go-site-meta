package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/siteMeta", func(c *gin.Context) {
		rawWebsiteURL := c.Query("url")
		websiteURL, sanitisationError := SanitiseURL(rawWebsiteURL)
		if sanitisationError != nil {
			c.JSON(400, gin.H{
				"error": "Something went wrong b",
			})
			return
		}

		imageURL, imageError := FindImageURL(websiteURL.String())
		if imageError != nil {
			c.JSON(400, gin.H{
				"error": "Something went wrong c",
			})
			return
		}

		c.JSON(200, gin.H{
			"websiteImageURL": imageURL,
		})
	})
	r.Run()
}
