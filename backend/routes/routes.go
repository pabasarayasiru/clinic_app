package routes

import (
	"clinic_app/controllers"
	"clinic_app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	controllers.Init(db)

	// AUTH
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	auth.POST("/refresh", controllers.RefreshToken)

	// PROTECTED ROUTES
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// VISITS
		api.POST("/visits", controllers.CreateVisit)
		api.GET("/visits/:id", controllers.GetVisit)
		api.GET("/visits", controllers.GetDoctorVisits)

		// CLINIC NOTES
		api.POST("/notes", controllers.AddClinicNote)
		api.GET("/notes/:visit_id", controllers.GetClinicNotes)

		// DRUGS
		api.POST("/drugs", controllers.AddDrug)
		api.GET("/drugs/:visit_id", controllers.GetDrugs)

		// LAB TESTS
		api.POST("/labs", controllers.AddLabTest)
		api.GET("/labs/:visit_id", controllers.GetLabTests)

		// BILLING
		api.POST("/billing", controllers.AddBilling)
		api.GET("/billing/:visit_id", controllers.GetBilling)

		// AI LOGS
		api.POST("/ai-logs", controllers.AddAIParsingLog)
		api.GET("/ai-logs/:visit_id", controllers.GetAIParsingLogs)
	}
}
