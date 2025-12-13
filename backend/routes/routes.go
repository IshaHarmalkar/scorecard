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
		
		api.POST("/scorecards", controllers.CreateScorecard)
		api.GET("/scorecards/:scorecardId", controllers.GetScorecard)
		api.GET("/scorecards", controllers.GetAllScorecards)
		api.PUT("/scorecards/:scorecardId", controllers.UpdateScorecard)
		api.DELETE("/scorecards/:scorecardId", controllers.DeleteScorecard)

			
		api.POST("/sections", controllers.CreateSection)
		api.GET("/scorecards/:scorecardId/sections/:sectionId", controllers.GetSection)
		api.GET("/sections", controllers.GetAllSections)
		api.PUT("/sections/:id", controllers.UpdateSection)
		api.DELETE("/sections/:id", controllers.DeleteSection)
	
	}
}