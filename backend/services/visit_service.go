package services

import (
	"clinic_app/models"
	"clinic_app/repositories"
	"gorm.io/gorm"
)

func CreateVisit(db *gorm.DB, visit models.Visit) error {
	return repositories.CreateVisit(db, visit)
}

func GetVisitByID(db *gorm.DB, id uint) (models.Visit, error) {
	return repositories.GetVisitByID(db, id)
}

func GetVisitsByDoctor(db *gorm.DB, doctorID uint) ([]models.Visit, error) {
	return repositories.GetVisitsByDoctor(db, doctorID)
}
