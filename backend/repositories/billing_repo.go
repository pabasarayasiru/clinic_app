package repositories

import (
	"clinic_app/models"
	"gorm.io/gorm"
)

func CreateBilling(db *gorm.DB, billing models.Billing) error {
	return db.Create(&billing).Error
}

func GetBillingByVisit(db *gorm.DB, visitID uint) (models.Billing, error) {
	var bill models.Billing
	err := db.Where("visit_id = ?", visitID).First(&bill).Error
	return bill, err
}
