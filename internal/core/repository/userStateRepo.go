package repository

import "langbrv/internal/core/model"

type UserStateRepo interface {
	Set(s *model.UserState) error
	Get(chatID int64) (*model.UserState, error)
}
