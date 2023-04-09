package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, world!")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/news-api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "News-API",
		})
	})

	router.Run(":8080")
}
