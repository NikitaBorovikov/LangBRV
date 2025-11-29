package repository

import "langbrv/internal/core/model"

type UserRepo interface {
	Create(u *model.User) (string, error)
	GetByID(userID string) (*model.User, error)
}
