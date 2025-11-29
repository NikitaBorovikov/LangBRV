package usecases

import "langbrv/internal/core/repository"

type UseCases struct {
	UserUC *UserUC
	WordUC *WordUC
}

func NewUseCases(ur repository.UserRepo, wr repository.WordRepo) *UseCases {
	return &UseCases{
		UserUC: NewUserUC(ur),
		WordUC: NewWordUC(wr),
	}
}
