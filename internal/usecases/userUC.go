package usecases

import (
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"time"
)

type UserUC struct {
	UserRepo repository.UserRepo
}

func NewUserUC(ur repository.UserRepo) *UserUC {
	return &UserUC{
		UserRepo: ur,
	}
}

func (uc *UserUC) CreateOrUpdate(user *model.User) error {
	user.CreatedAt = time.Now().UTC()
	return uc.UserRepo.CreateOrUpdate(user)
}

func (uc *UserUC) GetByID(userID string) (*model.User, error) {
	return nil, nil
}
