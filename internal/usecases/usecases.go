package usecases

import (
	repo "langbrv/internal/infrastucture/repository"
)

type UseCases struct {
	UserUC           *UserUC
	WordUC           *WordUC
	UserStateUC      *UserStateUC
	DictionaryPageUC *DictionaryPageUC
	RemindCardUC     *RemindCardUC
}

func NewUseCases(r *repo.Repository) *UseCases {
	return &UseCases{
		UserUC:           NewUserUC(r.UserRepo),
		WordUC:           NewWordUC(r.WordRepo),
		UserStateUC:      NewUserStateUC(r.UserStateRepo),
		DictionaryPageUC: NewDictionaryPageUC(r.DictionaryPageRepo, r.WordRepo),
		RemindCardUC:     NewRemindCardUC(r.RemindCardRepo, r.WordRepo),
	}
}
