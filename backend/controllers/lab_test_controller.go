package controllers

import (
	"clinic_app/models"
	"clinic_app/services"

	"github.com/gin-gonic/gin"
)

func AddLabTest(c *gin.Context) {
	var lab models.LabTest

	if err := c.ShouldBindJSON(&lab); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	err := services.AddLabTest(DB, lab)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Lab test added"})
}

func GetLabTests(c *gin.Context) {
	visitID := c.Param("visit_id")

	labs, _ := services.GetLabTests(DB, toUint(visitID))
	c.JSON(200, labs)
}
