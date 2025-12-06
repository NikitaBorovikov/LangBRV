package repository

import (
	inmemory "langbrv/internal/infrastucture/repository/inMemory"
	"langbrv/internal/infrastucture/repository/postgres"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepo           *postgres.UserRepo
	WordRepo           *postgres.WordRepo
	UserStateRepo      *inmemory.UserStateRepo
	DictionaryPageRepo *inmemory.DictionaryPageRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo:           postgres.NewUserRepo(db),
		WordRepo:           postgres.NewWordRepo(db),
		UserStateRepo:      inmemory.NewUserStateRepo(),
		DictionaryPageRepo: inmemory.NewDictionaryPageRepo(),
	}
}
