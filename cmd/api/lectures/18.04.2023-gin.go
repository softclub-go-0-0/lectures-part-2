package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!",
		})
	})

	r.GET("/users", GetUsers)

	r.Run(":6767")
}

func GetUsers(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "I'm on vacation!",
	})
}