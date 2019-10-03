package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/klogo/message"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/", requestHandler)
	return r
}

func TestRestHandler(t *testing.T) {

	q := &message.LogoRequest{
		ID:       "test",
		ImageURL: "https://storage.googleapis.com/kdemo-logos/youtube.jpg",
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
