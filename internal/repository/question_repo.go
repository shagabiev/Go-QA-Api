package repository

import (
	"github.com/shagabiev/Go-QA-Api/internal/models"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Create(q *models.Question) error {
	return r.db.Create(q).Error
}

func (r *QuestionRepository) FindAll() ([]models.Question, error) {
	var questions []models.Question
	err := r.db.Where("deleted_at IS NULL").Find(&questions).Error
	return questions, err
}

func (r *QuestionRepository) FindByID(id uint) (*models.Question, error) {
	var q models.Question
	err := r.db.Where("deleted_at IS NULL").First(&q, id).Error
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *QuestionRepository) FindByIDWithAnswers(id uint) (*models.Question, error) {
	var q models.Question
	err := r.db.
		Where("deleted_at IS NULL").
		Preload("Answers", "deleted_at IS NULL").
		First(&q, id).Error
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (r *QuestionRepository) Delete(id uint) error {
	result := r.db.
		Where("deleted_at IS NULL").
		Delete(&models.Question{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
