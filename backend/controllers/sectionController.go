package controllers

import (
	"fmt"
	"net/http"
	"scorecard/models"
	"scorecard/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sectionService = services.SectionService{}

func CreateSection(c *gin.Context) {

	var input models.Section

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("sectionController recieves input: ", input)

	if err := sectionService.CreateSection(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Section created"})

}


func GetSection(c *gin.Context) {

	scorecardId, err := strconv.Atoi(c.Param("scorecardId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"scorecard id must be an integer"})
		return
	}

	sectionId, err := strconv.Atoi(c.Param("sectionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"section id must be an integer"})
		return
	}
	
	

	section, err := sectionService.GetSection(scorecardId, sectionId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Section not found",
				"errorMsg": err.Error(),})
		return
	}

	c.JSON(http.StatusOK, section)

}

func GetAllSections(c *gin.Context){
	scorecards, err := sectionService.GetAllSections()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sections not found",
				"errorMsg": err.Error(),})
		return
	}

	c.JSON(http.StatusOK, scorecards)
}

func UpdateSection(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var input models.Section

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"id must be an integer"})
		return
	}
	input.Id = id

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	if err =  sectionService.UpdateSection(input); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update section", "errorMsg": err.Error(),})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Section updated"})

	
}


func DeleteSection(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"id must be an integer"})
		return
	}

	if err := sectionService.DeleteSection(id); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to delete", "errMsg": err.Error(),})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"Section deleted"})
	
}