package controllers

import (
	"clinic_app/models"
	"clinic_app/services"
	"clinic_app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}

func Register(c *gin.Context) {
	var doctor models.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := services.RegisterDoctor(DB, doctor)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Registered"})
}

func Login(c *gin.Context) {
	var input models.Doctor

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	doctor, err := services.LoginDoctor(DB, input.Email, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(doctor.ID)

	c.JSON(200, gin.H{
		"token": token,
	})
}
