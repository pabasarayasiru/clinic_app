package services

import (
	"clinic_app/models"
	"clinic_app/repositories"
	"gorm.io/gorm"
)

func AddAIParsingLog(db *gorm.DB, log models.AIParsingLog) error {
	return repositories.CreateAIParsingLog(db, log)
}

func GetAIParsingLogs(db *gorm.DB, visitID uint) ([]models.AIParsingLog, error) {
	return repositories.GetAIParsingLogsByVisit(db, visitID)
}
