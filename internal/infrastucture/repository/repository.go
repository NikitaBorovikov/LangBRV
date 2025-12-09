package repository

import (
	inmemory "langbrv/internal/infrastucture/repository/inMemory"
	"langbrv/internal/infrastucture/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepo           *postgres.UserRepo
	WordRepo           *postgres.WordRepo
	UserStateRepo      *inmemory.UserStateRepo
	DictionaryPageRepo *inmemory.DictionaryPageRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepo:           postgres.NewUserRepo(db),
		WordRepo:           postgres.NewWordRepo(db),
		UserStateRepo:      inmemory.NewUserStateRepo(),
		DictionaryPageRepo: inmemory.NewDictionaryPageRepo(),
	}
}
