package repository

import "langbrv/internal/core/model"

type RemindCardRepo interface {
	Save(card *model.RemindCard) error
	Get(userID int64) (*model.RemindCard, error)
}
