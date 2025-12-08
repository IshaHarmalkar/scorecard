package controllers

import (
	"net/http"
	"scorecard/models"
	"scorecard/services"

	"github.com/gin-gonic/gin"
)

var userService = services.UserService{}

func CreateUser(c *gin.Context) {

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userService.CreateUser(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
	

	

}