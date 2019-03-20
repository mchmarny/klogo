package main

import (
	"log"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"


)

const (
	defaultPort      = "8080"
	portVariableName = "PORT"
)

func main() {

	// router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// root & health
	r.GET("/", defaultHandler)
	r.POST("/", requestHandler)
	r.GET("/health", healthHandler)

	// port
	port := os.Getenv(portVariableName)
	if port == "" {
		port = defaultPort
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server starting: %s \n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}

}
