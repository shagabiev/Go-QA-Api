package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Answer struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	QuestionID uint           `gorm:"not null;index" json:"question_id"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	Text       string         `gorm:"not null" json:"text"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
