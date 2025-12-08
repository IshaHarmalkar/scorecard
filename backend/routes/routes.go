package routes

import (
	"scorecard/controllers"

	"github.com/gin-gonic/gin"
)



func SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/albums", controllers.GetAlbums)
		api.POST("/users", controllers.CreateUser)
	
	}
}