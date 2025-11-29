package repository

import "langbrv/internal/core/model"

type UserRepo interface {
	Create(u *model.User) (string, error)
	GetByID(userID string) (*model.User, error)
}

type UserStateRepo interface {
	Set(s *model.UserState) error
	Get(userID int64) (*model.UserState, error)
}
