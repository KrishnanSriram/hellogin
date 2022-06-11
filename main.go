package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"Hello from Golang-GIN",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.Run(":80")
}