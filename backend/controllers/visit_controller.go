package controllers

import (
	"clinic_app/models"
	"clinic_app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateVisit(c *gin.Context) {
	var visit models.Visit

	if err := c.ShouldBindJSON(&visit); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := services.CreateVisit(DB, visit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Visit created"})
}

func GetVisit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	visit, err := services.GetVisitByID(DB, uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	c.JSON(200, visit)
}

func GetDoctorVisits(c *gin.Context) {
	doctorID := c.GetUint("user_id")

	visits, _ := services.GetVisitsByDoctor(DB, doctorID)
	c.JSON(200, visits)
}
