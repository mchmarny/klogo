package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"log"
	"encoding/json"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/", requestHandler)
	return r
}

func TestRestHandler(t *testing.T) {

	q := &ServiceRequest{
		ID: "test",
		ImageURL: "https://storage.googleapis.com/kdemo-logos/0.png",
	}

	b, err := json.Marshal(q)
	assert.Nil(t, err)

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Printf(w.Body.String())
}
