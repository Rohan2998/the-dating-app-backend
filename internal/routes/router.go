package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// setting up router for gin
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")

	// v1 user routes
	UserRoutes(v1)

	return r
}
