package usecases

import "langbrv/internal/core/repository"

type UseCases struct {
	UserUC      *UserUC
	WordUC      *WordUC
	UserStateUC *UserStateUC
}

func NewUseCases(ur repository.UserRepo, wr repository.WordRepo, sr repository.UserStateRepo) *UseCases {
	return &UseCases{
		UserUC:      NewUserUC(ur),
		WordUC:      NewWordUC(wr),
		UserStateUC: NewUserStateUC(sr),
	}
}
