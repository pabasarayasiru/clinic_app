package models

import (
	"time"

	"gorm.io/gorm"
)

type DrugPrescription struct {
	ID        uint   `gorm:"primaryKey"`
	VisitID   uint   `gorm:"not null"`
	DrugName  string `gorm:"not null"`
	Dosage    string
	Frequency string
	Duration  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
