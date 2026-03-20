package middleware

import (
	"net/http"
	"strings"

	"clinic_app/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Format: Bearer TOKEN
		tokenString := strings.Split(authHeader, " ")[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// Set user info in context
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("role", claims["role"])

		c.Next()
	}
}
