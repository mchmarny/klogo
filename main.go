package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"

	ev "github.com/mchmarny/gcputil/env"
	pj "github.com/mchmarny/gcputil/project"
)

var (
	logger    = log.New(os.Stdout, "", 0)
	projectID = pj.GetIDOrFail()
	port      = ev.MustGetEnvVar("PORT", "8080")
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", defaultHandler)
	r.POST("/", requestHandler)
	r.GET("/health", healthHandler)
	r.NoRoute(defaultHandler)

	hostPort := net.JoinHostPort("0.0.0.0", port)
	logger.Printf("Starting server: %v", hostPort)
	if err := r.Run(hostPort); err != nil {
		logger.Fatal(err)
	}

}
