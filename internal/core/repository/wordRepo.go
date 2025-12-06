package repository

import "langbrv/internal/core/model"

type WordRepo interface {
	Add(word *model.Word) (string, error)
	Update(word *model.Word) error
	Delete(userID int64, word string) error
	GetDictionaryWordsByPage(userID, pageNum int64) ([]model.Word, error)
	GetRemindList(userID int64) ([]model.Word, error)
	GetAmountOfWords(userID int64) (int64, error)
	FindByUserAndWord(userID int64, word string) (*model.Word, error)
}
