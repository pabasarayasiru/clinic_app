package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateVisit(db *gorm.DB, visit models.Visit) error {
	return db.Create(&visit).Error
}

func GetVisitByID(db *gorm.DB, id uint) (models.Visit, error) {
	var visit models.Visit
	err := db.Preload("ClinicNotes").
		Preload("DrugPrescriptions").
		Preload("LabTests").
		Preload("Billing").
		First(&visit, id).Error
	return visit, err
}

func GetVisitsByDoctor(db *gorm.DB, doctorID uint) ([]models.Visit, error) {
	var visits []models.Visit
	err := db.Where("doctor_id = ?", doctorID).
		Preload("ClinicNotes").
		Preload("DrugPrescriptions").
		Preload("LabTests").
		Preload("Billing").
		Find(&visits).Error
	return visits, err
}
