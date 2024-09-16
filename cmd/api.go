package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func get_port() string {
	port := ":8082"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	r := gin.Default()
	r.GET("/api/ping", getPing)
	r.GET("/api/testlong", getTestLong)
	r.GET("/api/testshort", getTestShort)
	r.Run(get_port())
}
