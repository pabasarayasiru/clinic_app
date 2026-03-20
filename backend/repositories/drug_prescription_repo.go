package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateDrugPrescription(db *gorm.DB, drug models.DrugPrescription) error {
	return db.Create(&drug).Error
}

func GetDrugsByVisit(db *gorm.DB, visitID uint) ([]models.DrugPrescription, error) {
	var drugs []models.DrugPrescription
	err := db.Where("visit_id = ?", visitID).Find(&drugs).Error
	return drugs, err
}
