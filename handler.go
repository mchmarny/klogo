package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/klogo/message"
)

const (
	tokenKey = "token"
)

// requestHandler handles posted queries
func requestHandler(c *gin.Context) {

	// bind
	var req message.LogoRequest
	err := c.BindJSON(&req)
	if err != nil {
		logger.Printf("Error while binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	logger.Printf("Request: %v", req)

	// process
	desc, err := getLogoFromURL(req.ImageURL)
	if err != nil {
		logger.Printf("Error while getting logo: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error processing request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	resp := &message.LogoResponse{
		Request:     req,
		Description: desc,
	}

	c.JSON(http.StatusOK, resp)

}

// defaultHandler handles the root query
func defaultHandler(c *gin.Context) {

	logger.Println("Default handler...")

	// sample request
	req := &message.LogoRequest{
		ID:       "sample-id",
		ImageURL: "https://storage.googleapis.com/kdemo-logos/0.png",
	}

	c.JSON(http.StatusOK, req)

}

func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
