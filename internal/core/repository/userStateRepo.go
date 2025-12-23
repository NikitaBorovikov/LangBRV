package repository

import "langbrv/internal/core/model"

type UserStateRepo interface {
	Save(s *model.UserState) error
	Get(userID int64) (*model.UserState, error)
}
