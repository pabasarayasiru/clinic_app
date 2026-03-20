package controllers

import (
	"clinic_app/models"
	"clinic_app/services"

	"github.com/gin-gonic/gin"
)

func AddBilling(c *gin.Context) {
	var bill models.Billing

	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	err := services.AddBilling(DB, bill)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Billing added"})
}

func GetBilling(c *gin.Context) {
	visitID := c.Param("visit_id")

	bill, _ := services.GetBilling(DB, toUint(visitID))
	c.JSON(200, bill)
}
