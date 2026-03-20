package models

import (
	"time"

	"gorm.io/gorm"
)

type AIParsingLog struct {
	ID        uint   `gorm:"primaryKey"`
	VisitID   uint   `gorm:"not null"`
	RawInput  string `gorm:"type:text;not null"`
	Drugs     string `gorm:"type:text"`
	LabTests  string `gorm:"type:text"`
	Notes     string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
