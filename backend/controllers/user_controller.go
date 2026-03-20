package controllers

import (
	"clinic_app/models"
	"clinic_app/services"
	"clinic_app/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := services.RegisterUser(DB, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Registered"})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	user, err := services.LoginUser(DB, input.Email, input.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, refreshToken := utils.GenerateTokenPair(user.ID, user.Role)

	c.JSON(200, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	token, err := utils.ValidateToken(body.RefreshToken)
	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	userID := uint(claims["user_id"].(float64))
	role := "doctor" // or fetch from DB if needed

	newAccess, newRefresh := utils.GenerateTokenPair(userID, role)

	c.JSON(200, gin.H{
		"token":         newAccess,
		"refresh_token": newRefresh,
	})
}
