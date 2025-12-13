package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"strings"
)

const (
	wordsPerPage     int64 = 5
	expectedPageSize int   = 700 // примерный размер одной страницы словаря в байтах
)

type DictionaryPageUC struct {
	DictionaryPageRepo repository.DictionaryPageRepo
	WordRepo           repository.WordRepo
}

func NewDictionaryPageUC(pr repository.DictionaryPageRepo, wr repository.WordRepo) *DictionaryPageUC {
	return &DictionaryPageUC{
		DictionaryPageRepo: pr,
		WordRepo:           wr,
	}
}

func (uc *DictionaryPageUC) FormatPage(pageInfo *model.DictionaryPage) (string, error) {
	words, err := uc.WordRepo.GetDictionaryWordsByPage(pageInfo.UserID, pageInfo.CurrentPage, wordsPerPage)
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

func (uc *DictionaryPageUC) Save(page *model.DictionaryPage) error {
	return uc.DictionaryPageRepo.Save(page)
}

func (uc *DictionaryPageUC) Get(userID int64) (*model.DictionaryPage, error) {
	return uc.DictionaryPageRepo.Get(userID)
}
