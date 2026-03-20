package controllers

import (
	"clinic_app/models"
	"clinic_app/services"

	"github.com/gin-gonic/gin"
)

func AddClinicNote(c *gin.Context) {
	var note models.ClinicNote

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	err := services.AddClinicNote(DB, note)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Note added"})
}

func GetClinicNotes(c *gin.Context) {
	visitID := c.Param("visit_id")

	notes, _ := services.GetClinicNotes(DB, toUint(visitID))
	c.JSON(200, notes)
}
