package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetUserByID(db *gorm.DB, id uint) (models.User, error) {
	var user models.User
	err := db.First(&user, id).Error
	return user, err
}
