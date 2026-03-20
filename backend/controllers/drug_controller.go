package controllers

import (
	"clinic_app/models"
	"clinic_app/services"

	"github.com/gin-gonic/gin"
)

func AddDrug(c *gin.Context) {
	var drug models.DrugPrescription

	if err := c.ShouldBindJSON(&drug); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	err := services.AddDrugPrescription(DB, drug)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Drug added"})
}

func GetDrugs(c *gin.Context) {
	visitID := c.Param("visit_id")

	drugs, _ := services.GetDrugs(DB, toUint(visitID))
	c.JSON(200, drugs)
}
