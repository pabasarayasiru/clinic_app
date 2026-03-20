package models

import (
	"time"

	"gorm.io/gorm"
)

type LabTest struct {
	ID        uint   `gorm:"primaryKey"`
	VisitID   uint   `gorm:"not null"`
	TestName  string `gorm:"not null"`
	Status    string `gorm:"type:enum('ordered','completed','canceled');default:'ordered'"`
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
