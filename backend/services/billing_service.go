package services

import (
	"clinic_app/models"
	"clinic_app/repositories"
	"gorm.io/gorm"
)

func AddBilling(db *gorm.DB, bill models.Billing) error {
	return repositories.CreateBilling(db, bill)
}

func GetBilling(db *gorm.DB, visitID uint) (models.Billing, error) {
	return repositories.GetBillingByVisit(db, visitID)
}
