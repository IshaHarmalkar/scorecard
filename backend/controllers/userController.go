package controllers

import (
	"net/http"
	"scorecard/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}




}