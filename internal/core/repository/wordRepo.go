package repository

import "langbrv/internal/core/model"

type WordRepo interface {
	Add(word *model.Word) (string, error)
	GetDictionaryWordsByPage(userID int64, pageNum int) ([]model.Word, error)
	GetAmountOfWords(userID int64) (int64, error)
	FindByUserAndWord(userID int64, word string) (*model.Word, error)
	GetRemindList(userID int64) ([]model.Word, error)
	Update(word *model.Word) error
	DeleteWord(userID int64, word string) error
}
