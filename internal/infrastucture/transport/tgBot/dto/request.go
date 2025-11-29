package dto

import (
	"fmt"
	"langbrv/internal/core/model"
	"strings"
	"time"
)

type AddWordRequest struct {
	UserID int64
	Msg    string
}

func NewAddWordRequest(userID int64, msg string) *AddWordRequest {
	return &AddWordRequest{
		UserID: userID,
		Msg:    msg,
	}
}

func (r *AddWordRequest) ToDomainWord() (*model.Word, error) {
	record := strings.Split(r.Msg, "-")

	if r.UserID == 0 {
		return nil, fmt.Errorf("невалидный userID")
	}

	if len(record) != 2 {
		return nil, fmt.Errorf("слова должны быть разделены черточкой")
	}

	if len(record[0]) >= 255 || len(record[1]) >= 255 {
		return nil, fmt.Errorf("слова слишком длинные")
	}

	word := &model.Word{
		UserID:      r.UserID,
		UserWord:    strings.TrimSpace(record[0]),
		Translation: strings.TrimSpace(record[1]),
		CreatedAt:   time.Now(),
	}
	return word, nil
}
