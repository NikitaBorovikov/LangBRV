package dto

import (
	"fmt"
	"langbrv/internal/core/model"
	"strings"
	"time"
)

type RegistrationRequest struct {
	UserID   int64
	Username string
}

func NewRegistrationRequest(userID int64, username string) *RegistrationRequest {
	return &RegistrationRequest{
		UserID:   userID,
		Username: username,
	}
}

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
		Original:    strings.TrimSpace(record[0]),
		Translation: strings.TrimSpace(record[1]),
		LastSeen:    time.Now(),
	}
	return word, nil
}

func (r *RegistrationRequest) ToDomainUser() *model.User {
	return &model.User{
		ID:       r.UserID,
		Username: r.Username,
	}
}
