package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateClinicNote(db *gorm.DB, note models.ClinicNote) error {
	return db.Create(&note).Error
}

func GetClinicNotesByVisit(db *gorm.DB, visitID uint) ([]models.ClinicNote, error) {
	var notes []models.ClinicNote
	err := db.Where("visit_id = ?", visitID).Find(&notes).Error
	return notes, err
}
