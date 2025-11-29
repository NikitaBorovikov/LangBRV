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

func (r *WordRepo) Add(w *model.Word) (string, error) {
	return "", nil
}

func (r *WordRepo) GetAll(userID string) ([]model.Word, error) {
	return nil, nil
}
