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
		api.GET("/users/:id", controllers.GetUser)
		api.GET("/users", controllers.GetAllUsers)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)


		
		api.POST("/scorecard", controllers.CreateScorecard)
		api.GET("/scorecard/:id", controllers.GetScorecard)
		api.GET("/scorecard", controllers.GetAllScorecards)
		api.PUT("/scorecard/:id", controllers.UpdateScorecard)
		api.DELETE("/scorecard/:id", controllers.DeleteScorecard)
	
	}
}