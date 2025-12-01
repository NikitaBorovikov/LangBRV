package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
)

type WordUC struct {
	WordRepo repository.WordRepo
}

func NewWordUC(wr repository.WordRepo) *WordUC {
	return &WordUC{
		WordRepo: wr,
	}
}

func (uc *WordUC) Add(word *model.Word) (string, error) {
	wordID, err := uc.WordRepo.Add(word)
	return wordID, err
}

func (uc *WordUC) GetAll(userID int64) ([]model.Word, error) {
	words, err := uc.WordRepo.GetAll(userID)
	if err != nil {
		return nil, err
	}
	return words, err
}

func (uc *WordUC) FormatDictionary(words []model.Word) (string, error) {
	if len(words) == 0 {
		return "", apperrors.ErrNoWordsInDictionary
	}

	dictionary := "Твой словарь:\n"

	for idx, word := range words {
		dictionary += fmt.Sprintf("%d. %s - %s\n", idx+1, word.Original, word.Translation)
	}
	return dictionary, nil
}
