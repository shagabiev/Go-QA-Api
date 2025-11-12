package repository

import (
	"github.com/shagabiev/Go-QA-Api/internal/models"
	"gorm.io/gorm"
)

type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

func (r *AnswerRepository) Create(a *models.Answer) error {
	return r.db.Create(a).Error
}

func (r *AnswerRepository) FindByID(id uint) (*models.Answer, error) {
	var a models.Answer
	err := r.db.Where("deleted_at IS NULL").First(&a, id).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AnswerRepository) Delete(id uint) error {
	result := r.db.Where("id = ? AND deleted_at IS NULL").Delete(&models.Answer{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
