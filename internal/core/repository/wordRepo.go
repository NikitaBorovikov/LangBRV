package repository

import "langbrv/internal/core/model"

type WordRepo interface {
	Add(w *model.Word) (string, error)
	GetAll(userID int64) ([]model.Word, error)
	FindByUserAndWord(userID int64, word string) (*model.Word, error)
}
