package dto

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"strings"
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
	record := strings.Split(strings.ToLower(r.Msg), "-")
	if len(record) != 2 {
		return nil, apperrors.ErrIncorrectFormat
	}

	original := strings.TrimSpace(record[0])
	translate := strings.TrimSpace(record[1])

	if err := ValidateWord(original); err != nil {
		return nil, err
	}

	if err := ValidateWord(translate); err != nil {
		return nil, err
	}

	word := &model.Word{
		UserID:      r.UserID,
		Original:    original,
		Translation: translate,
	}
	return word, nil
}
