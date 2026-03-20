package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateLabTest(db *gorm.DB, lab models.LabTest) error {
	return db.Create(&lab).Error
}

func GetLabTestsByVisit(db *gorm.DB, visitID uint) ([]models.LabTest, error) {
	var labs []models.LabTest
	err := db.Where("visit_id = ?", visitID).Find(&labs).Error
	return labs, err
}
