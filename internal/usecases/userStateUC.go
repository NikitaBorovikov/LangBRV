package usecases

import (
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
)

type UserStateUC struct {
	UserStateRepo repository.UserStateRepo
}

func NewUserStateUC(sr repository.UserStateRepo) *UserStateUC {
	return &UserStateUC{UserStateRepo: sr}
}

func (uc *UserStateUC) Save(s *model.UserState) error {
	return uc.UserStateRepo.Save(s)
}

func (uc *UserStateUC) Get(chatID int64) (*model.UserState, error) {
	return uc.UserStateRepo.Get(chatID)
}
