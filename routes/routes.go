package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to arash-website API!",
		})
	})
}
