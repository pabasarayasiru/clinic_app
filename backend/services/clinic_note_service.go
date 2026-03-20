package services

import (
	"clinic_app/models"
	"clinic_app/repositories"
	"gorm.io/gorm"
)

func AddClinicNote(db *gorm.DB, note models.ClinicNote) error {
	return repositories.CreateClinicNote(db, note)
}

func GetClinicNotes(db *gorm.DB, visitID uint) ([]models.ClinicNote, error) {
	return repositories.GetClinicNotesByVisit(db, visitID)
}
