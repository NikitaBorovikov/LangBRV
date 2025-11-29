package postgres

import (
	"langbrv/internal/core/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(u *model.User) (string, error) {
	return "", nil
}

func (r *UserRepo) GetByID(userID string) (*model.User, error) {
	return nil, nil
}
