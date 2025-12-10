package controllers

import (
	"net/http"
	"scorecard/models"
	"scorecard/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var scorecardService = services.ScorecardService{}

func CreateScorecard(c *gin.Context) {

	var input models.Scorecard

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := scorecardService.CreateScorecard(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Scorecard created"})

}


func GetScorecard(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"id must be an integer"})
		return

	}
	

	user, err := userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetAllScorecards(c *gin.Context){
	users, err := userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Could not fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateScorecard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var input models.User

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"id must be an integer"})
		return
	}
	input.Id = id

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	if err =  userService.UpdateUser(input); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})

	
}


func DeleteScorecard(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"id must be an integer"})
		return
	}

	if err := userService.DeleteUser(id); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"User deleted"})
	
}