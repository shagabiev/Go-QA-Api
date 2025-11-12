package models

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Text      string         `gorm:"not null" json:"text"`
	CreatedAt time.Time      `json:"created_at"`
	Answers   []Answer       `gorm:"constraint:OnDelete:CASCADE" json:"answers,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
