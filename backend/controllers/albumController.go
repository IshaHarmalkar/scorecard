package controllers

import (
	"scorecard/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums = []models.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},

	}
	c.IndentedJSON(http.StatusOK, albums)
}