package repository

import "langbrv/internal/core/model"

type DictionaryPageRepo interface {
	Save(page *model.DictionaryPage) error
	Get(userID int64) (*model.DictionaryPage, error)
}
