package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"relay/services"
)

func CreateShortUrl(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	shortUrl, err := services.GetShortUrl(request.URL)
	if err != nil {
		log.Printf("Error generating short URL for %s: %v", request.URL, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"short_url": shortUrl})
}
