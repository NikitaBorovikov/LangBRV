package usecases

import (
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
	return words, err
}
