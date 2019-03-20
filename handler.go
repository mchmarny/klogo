package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	tokenKey = "token"
)

// requestHandler handles posted queries
func requestHandler(c *gin.Context) {


	// bind
	var req ServiceRequest
    err := c.BindJSON(&req)
    if err != nil {
		log.Printf("Error while binding request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	log.Printf("Request: %v", req)


	// process
	desc, err := getLogoFromUrl(req.ImageURL)
	if err != nil {
		log.Printf("Error while getting logo: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error processing request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	resp := &ServiceResponse{
		ID: req.ID,
		ImageURL: req.ImageURL,
		Description: desc,
	}


	c.JSON(http.StatusOK, resp)

}



// defaultHandler handles the root query
func defaultHandler(c *gin.Context) {

	// sample request
	req := &ServiceRequest{
		ID: "sample-id",
		ImageURL: "https://storage.googleapis.com/kdemo-logos/0.png",
	}


	c.JSON(http.StatusOK, req)

}


func healthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}