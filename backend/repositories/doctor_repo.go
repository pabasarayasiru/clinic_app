package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateDoctor(db *gorm.DB, doctor models.Doctor) error {
	return db.Create(&doctor).Error
}

func GetDoctorByEmail(db *gorm.DB, email string) (models.Doctor, error) {
	var doctor models.Doctor
	err := db.Where("email = ?", email).First(&doctor).Error
	return doctor, err
}
