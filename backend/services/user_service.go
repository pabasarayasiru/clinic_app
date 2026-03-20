package services

import (
	"clinic_app/models"
	"clinic_app/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, user models.User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	return repositories.CreateUser(db, user)
}

func LoginUser(db *gorm.DB, email, password string) (models.User, error) {
	user, err := repositories.GetUserByEmail(db, email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
