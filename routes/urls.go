package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"relay/services"
)

func CreateShortUrl(c *gin.Context) {
	var longUrl string
	
	if c.ContentType() == "application/x-www-form-urlencoded" {
		longUrl = c.PostForm("url")
	} else {
		var request struct {
			URL string `json:"url" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			log.Printf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		longUrl = request.URL
	}

	if longUrl == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	shortCode, err := services.GetShortUrl(longUrl)
	if err != nil {
		log.Printf("Error generating short URL for %s: %v", longUrl, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	host := c.Request.Host
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	fullShortUrl := scheme + "://" + host + "/" + shortCode

	if strings.Contains(c.GetHeader("Accept"), "text/html") || c.ContentType() == "application/x-www-form-urlencoded" {
		c.HTML(http.StatusOK, "result.html", gin.H{"shortUrl": fullShortUrl})
	} else {
		c.JSON(http.StatusCreated, gin.H{"short_url": fullShortUrl})
	}
}

func Forward(c *gin.Context) {
	code := c.Param("code")

	longUrl, err := services.GetLongUrl(code)

	if longUrl == "" || err != nil {
		ServeErrorPage(c, "404 Not Found: The requested URL does not exist.")
		return
	}

	c.Redirect(http.StatusMovedPermanently, longUrl)
}
