package repository

import "langbrv/internal/core/model"

type UserRepo interface {
	CreateOrUpdate(u *model.User) error
	GetByID(userID string) (*model.User, error)
}
