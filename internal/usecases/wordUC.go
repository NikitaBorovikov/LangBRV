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

func (uc *WordUC) Add(w *model.Word) (string, error) {
	return "", nil
}

func (uc *WordUC) GetAll(userID string) ([]model.Word, error) {
	return nil, nil
}
