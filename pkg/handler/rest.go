package handler

import (
	"log"
	"net/http"

	// "github.com/mchmarny/kcm/pkg/processor"
	"github.com/mchmarny/kcm/pkg/common"

	"github.com/gin-gonic/gin"
)

const (
	tokenKey = "token"
)

// RestHandler handles queries from the REST interface
func RestHandler(c *gin.Context) {

	// token
	if !common.IsValidAccessToken(tokenKey, c.Query("token")) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid access token",
			"status":  http.StatusBadRequest,
		})
		return
	}

	query := c.Param("query")
	if query == "" {
		log.Println("nil id")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Required parameter missing: query",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// result, err := processor.Search(c.Request.Context(), query)
	// if err != nil {
	// 	log.Printf("error on search: %v", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": err,
	// 		"status":  http.StatusBadRequest,
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, result)

	c.JSON(http.StatusOK, "")

}
