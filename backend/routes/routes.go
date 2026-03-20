package routes

import (
	"clinic_app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	controllers.Init(db)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
