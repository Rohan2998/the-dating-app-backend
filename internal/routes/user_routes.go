package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Hello World"})
	})

	router.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Hello World"})
	})

	router.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Hello World"})
	})

	router.PUT("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Hello World"})
	})

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Hello World"})
	})
}
