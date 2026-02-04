package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"strings"
)

type RemindCardUC struct {
	WordRepo repository.WordRepo
}

func NewRemindCardUC(wr repository.WordRepo) *RemindCardUC {
	return &RemindCardUC{
		WordRepo: wr,
	}
}

func (uc *RemindCardUC) FormatClosedRemindCard(remindCards model.RemindSession) (string, error) {
	if len(remindCards.Words) == 0 {
		return "", apperrors.ErrNoWordsToRemind
	}

	currentWord := remindCards.Words[remindCards.CurrentCard-1]

	var sb strings.Builder
	sb.Grow(expectedPageSize)
	fmt.Fprintf(&sb, "🌀 <b>Повторение:</b> <i>%d/%d</i>\n\n", remindCards.CurrentCard, remindCards.TotalCards)
	fmt.Fprintf(&sb, "<b>%s - _________</b>", currentWord.Original)
	return sb.String(), nil
}

func (uc *RemindCardUC) FormatOpenedRemindCard(remindCards model.RemindSession) (string, error) {
	if len(remindCards.Words) == 0 {
		return "", apperrors.ErrNoWordsToRemind
	}

	currentWord := remindCards.Words[remindCards.CurrentCard-1]

	var sb strings.Builder
	sb.Grow(expectedPageSize)
	fmt.Fprintf(&sb, "🌀 <b>Повторение:</b> <i>%d/%d</i>\n\n", remindCards.CurrentCard, remindCards.TotalCards)
	fmt.Fprintf(&sb, "<b>%s - %s</b>\n\n", currentWord.Original, currentWord.Translation)

	// The first card shows the instructions.
	if remindCards.CurrentCard == 1 {
		fmt.Fprintf(&sb, "<i>👎 - помню плохо. 👍 - помню хорошо.</i>")
	}
	return sb.String(), nil
}

func (uc *RemindCardUC) FormatListOfRemindedWords(remindList []model.Word) (string, error) {
	if len(remindList) == 0 {
		return "", apperrors.ErrNoWordsToRemind
	}

	var sb strings.Builder
	sb.Grow(expectedPageSize)
	fmt.Fprintf(&sb, "🦦 <b>Вот слова, которые ты сегодня повторил:</b> \n\n")

	for _, word := range remindList {
		fmt.Fprintf(&sb, "• %s - %s\n", word.Original, word.Translation)
	}
	return sb.String(), nil
}
