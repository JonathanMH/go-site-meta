package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/siteMeta", func(c *gin.Context) {
		websiteURL, sanitisationError := SanitiseURL(c.Query("websiteURL"))
		if sanitisationError != nil {
			c.JSON(400, gin.H{
				"error": "Something went wrong",
			})
		}
		imageURL, imageError := FindImageURL(websiteURL.String())
		if imageError != nil {
			c.JSON(400, gin.H{
				"error": "Something went wrong",
			})
		}
		c.JSON(200, gin.H{
			"websiteImageURL": imageURL,
		})
	})
	r.Run()
}
