package usecases

import (
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
)

type UserUC struct {
	UserRepo repository.UserRepo
}

func NewUserUC(ur repository.UserRepo) *UserUC {
	return &UserUC{
		UserRepo: ur,
	}
}

func (uc *UserUC) Create(u *model.User) (string, error) {
	return "", nil
}

func (uc *UserUC) GetByID(userID string) (*model.User, error) {
	return nil, nil
}
