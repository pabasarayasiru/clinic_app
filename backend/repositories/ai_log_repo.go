package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateAIParsingLog(db *gorm.DB, log models.AIParsingLog) error {
	return db.Create(&log).Error
}

func GetAIParsingLogsByVisit(db *gorm.DB, visitID uint) ([]models.AIParsingLog, error) {
	var logs []models.AIParsingLog
	err := db.Where("visit_id = ?", visitID).Find(&logs).Error
	return logs, err
}
