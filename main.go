package main

import (
	"log"
	"net"
	"net/http"
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

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", defaultHandler)
	r.POST("/", requestHandler)
	r.GET("/health", healthHandler)

	hostPort := net.JoinHostPort("0.0.0.0", port)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		logger.Fatal(err)
	}

}
