package services

import (
	"clinic_app/models"
	"clinic_app/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterDoctor(db *gorm.DB, doctor models.Doctor) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(doctor.Password), 10)
	doctor.Password = string(hashed)

	return repositories.CreateDoctor(db, doctor)
}

func LoginDoctor(db *gorm.DB, email, password string) (models.Doctor, error) {
	doctor, err := repositories.GetDoctorByEmail(db, email)
	if err != nil {
		return doctor, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(password))
	return doctor, err
}
