package models

import (
	"time"

	"gorm.io/gorm"
)

type ClinicNote struct {
	ID        uint   `gorm:"primaryKey"`
	VisitID   uint   `gorm:"not null"`
	NoteText  string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
