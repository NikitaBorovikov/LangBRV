package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"strings"
)

const (
	wordsPerPage     int64 = 10
	expectedPageSize int   = 700 // The approximate size of one dictionary page in bytes
)

type DictionaryPageUC struct {
	WordRepo repository.WordRepo
}

func NewDictionaryPageUC(wr repository.WordRepo) *DictionaryPageUC {
	return &DictionaryPageUC{
		WordRepo: wr,
	}
}

func (uc *DictionaryPageUC) FormatPage(userID int64, pageInfo *model.DictionaryPage) (string, error) {
	words, err := uc.WordRepo.GetDictionaryWordsByPage(userID, pageInfo.CurrentPage, wordsPerPage)
	if err != nil {
		return "", err
	}

	if len(words) == 0 {
		return "", apperrors.ErrNoWordsInDictionary
	}
	pageInfo.Words = words

	var sb strings.Builder
	sb.Grow(expectedPageSize)

	sb.WriteString("📚 <b>Твой словарь:</b>")
	fmt.Fprintf(&sb, " <i>(cтраница %d/%d)</i>\n\n", pageInfo.CurrentPage, pageInfo.TotalPages)

	for _, word := range pageInfo.Words {
		fmt.Fprintf(&sb, "• %s - %s\n", word.Original, word.Translation)
	}
	return sb.String(), nil
}

func (uc *DictionaryPageUC) GetAmountOfPages(userID int64) (int64, error) {
	amountOfWords, err := uc.WordRepo.GetAmountOfWords(userID)
	if err != nil {
		return 0, err
	}

	if amountOfWords == 0 {
		return 0, apperrors.ErrNoWordsInDictionary
	}

	totalPages := (amountOfWords + wordsPerPage - 1) / wordsPerPage
	return totalPages, nil
}
