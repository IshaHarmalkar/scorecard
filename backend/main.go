package main

import (
	"scorecard/database"
	"scorecard/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	database.Connect()
	r := gin.Default()
	routes.SetRoutes(r)
	r.Run(":8080")


	
}