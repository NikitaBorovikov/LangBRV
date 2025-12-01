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

func (r *UserRepo) CreateOrUpdate(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) GetByID(userID string) (*model.User, error) {
	var user model.User
	result := r.db.Where("id = ?", userID).Find(&user)
	return &user, result.Error
}
