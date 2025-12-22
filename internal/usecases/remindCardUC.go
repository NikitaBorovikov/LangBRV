package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"strings"
)

type RemindCardUC struct {
	RemindCardRepo repository.RemindCardRepo
	WordRepo       repository.WordRepo
}

func NewRemindCardUC(rc repository.RemindCardRepo, wr repository.WordRepo) *RemindCardUC {
	return &RemindCardUC{
		RemindCardRepo: rc,
		WordRepo:       wr,
	}
}

func (uc *RemindCardUC) FormatRemindCard(remindCards model.RemindCard) (string, error) {
	if len(remindCards.Words) == 0 {
		return "", apperrors.ErrNoWordsToRemind
	}

	currentWord := remindCards.Words[remindCards.CurrentCard-1]

	var sb strings.Builder
	sb.Grow(expectedPageSize)
	fmt.Fprintf(&sb, "🌀 <b>Повторение:</b> <i>%d/%d</i>\n\n", remindCards.CurrentCard, remindCards.TotalCards)
	fmt.Fprintf(&sb, "%s - %s", currentWord.Original, currentWord.Translation)
	return sb.String(), nil
}

func (uc *RemindCardUC) Save(card *model.RemindCard) error {
	return uc.RemindCardRepo.Save(card)
}

func (uc *RemindCardUC) Get(userID int64) (*model.RemindCard, error) {
	return uc.RemindCardRepo.Get(userID)
}
