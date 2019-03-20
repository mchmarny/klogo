package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mchmarny/kcm/pkg/handler"

	"github.com/gin-gonic/gin"


)

const (
	defaultPort      = "8080"
	portVariableName = "PORT"
)

func main() {

	log.Print("Initializing API server...")

	// router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// root & health
	r.GET("/", handler.RestHandler)
	r.GET("/health", defaultHandler)

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

func defaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
