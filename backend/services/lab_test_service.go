package services

import (
	"clinic_app/models"
	"clinic_app/repositories"
	"gorm.io/gorm"
)

func AddLabTest(db *gorm.DB, lab models.LabTest) error {
	return repositories.CreateLabTest(db, lab)
}

func GetLabTests(db *gorm.DB, visitID uint) ([]models.LabTest, error) {
	return repositories.GetLabTestsByVisit(db, visitID)
}
