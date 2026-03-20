package services

import (
	"clinic_app/models"
	"clinic_app/repositories"
	"gorm.io/gorm"
)

func AddDrugPrescription(db *gorm.DB, drug models.DrugPrescription) error {
	return repositories.CreateDrugPrescription(db, drug)
}

func GetDrugs(db *gorm.DB, visitID uint) ([]models.DrugPrescription, error) {
	return repositories.GetDrugsByVisit(db, visitID)
}
