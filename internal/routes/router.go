package routes

import "github.com/gin-gonic/gin"

// setting up router for gin
func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}
