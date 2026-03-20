package models

import (
	"time"

	"gorm.io/gorm"
)

type Billing struct {
	ID             uint `gorm:"primaryKey"`
	VisitID        uint `gorm:"unique;not null"`
	DrugCharges    float64
	LabTestCharges float64
	OtherCharges   float64
	TotalAmount    float64
	PaymentStatus  string `gorm:"type:enum('paid','unpaid');default:'unpaid'"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
