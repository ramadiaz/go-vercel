package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.New()

	router.GET("/api/hello", HelloHandler)
	router.POST("/api/post-test", PostTestHandler)

	router.ServeHTTP(w, r)
}
//

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Golang with Gin on Vercel!")
}

func PostTestHandler(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"received_message": json.Message,
	})
}
