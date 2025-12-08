package main

import (
	"scorecard/models"
	"scorecard/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	models.Connect()

    r := gin.Default()
	routes.SetRoutes(r)
	r.Run(":8080")





	
}