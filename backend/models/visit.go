package models

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	ID             uint   `gorm:"primaryKey"`
	DoctorID       uint   `gorm:"not null"`
	PatientName    string `gorm:"not null"`
	PatientAge     uint
	PatientGender  string
	PatientContact string
	VisitDate      time.Time `gorm:"not null"`
	NotesRawInput  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	// Relationships
	ClinicNotes       []ClinicNote       `gorm:"foreignKey:VisitID"`
	DrugPrescriptions []DrugPrescription `gorm:"foreignKey:VisitID"`
	LabTests          []LabTest          `gorm:"foreignKey:VisitID"`
	Billing           Billing            `gorm:"foreignKey:VisitID"`
}
