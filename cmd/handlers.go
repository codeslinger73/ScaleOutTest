package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTestShort(c *gin.Context) {
	s := c.Query("iterations")
	iterations, err := strconv.Atoi(s)
	if err != nil || iterations == 0 {
		iterations = 1500
	}
	processShort(iterations)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func getTestLong(c *gin.Context) {
	s := c.Query("iterations")
	iterations, err := strconv.Atoi(s)
	if err != nil || iterations == 0 {
		iterations = 6000
	}
	processLong(iterations)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Service is up",
	})
}
