package controllers

import (
	"clinic_app/models"
	"clinic_app/services"

	"github.com/gin-gonic/gin"
)

func AddAIParsingLog(c *gin.Context) {
	var log models.AIParsingLog

	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(400, gin.H{"error": "Invalid"})
		return
	}

	err := services.AddAIParsingLog(DB, log)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "AI log saved"})
}

func GetAIParsingLogs(c *gin.Context) {
	visitID := c.Param("visit_id")

	logs, _ := services.GetAIParsingLogs(DB, toUint(visitID))
	c.JSON(200, logs)
}
