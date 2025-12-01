package postgres

import (
	"langbrv/internal/core/model"

	"gorm.io/gorm"
)

type WordRepo struct {
	db *gorm.DB
}

func NewWordRepo(db *gorm.DB) *WordRepo {
	return &WordRepo{
		db: db,
	}
}

func (r *WordRepo) Add(word *model.Word) (string, error) {
	result := r.db.Create(word)
	return word.ID, result.Error
}

func (r *WordRepo) GetAll(userID string) ([]model.Word, error) {
	var words []model.Word
	result := r.db.Where("user_id = ?", userID).Find(&words)
	return words, result.Error
}
