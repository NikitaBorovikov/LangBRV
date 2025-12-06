package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"strings"
	"time"
)

const (
	wordsPerPage = 5
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
	// Проверяем, есть ли уже такое слово в словаре
	existingWord, err := uc.WordRepo.FindByUserAndWord(word.UserID, word.Original)
	if err != nil {
		return "", err
	}

	// Если слово уже есть, то просто обновляем его с новым LastSeen полем
	if existingWord != nil {
		existingWord.LastSeen = time.Now()
		err := uc.WordRepo.Update(existingWord)
		return existingWord.ID, err
	}
	// Если слова нет, то добавляем его
	wordID, err := uc.WordRepo.Add(word)
	if err != nil {
		return "", err
	}
	return wordID, nil
}

func (uc *WordUC) Delete(userID int64, word string) error {
	if err := dto.ValidateWord(word); err != nil {
		return err
	}

	if err := uc.WordRepo.Delete(userID, word); err != nil {
		return err
	}
	return nil
}

func (uc *WordUC) GetAmountOfPages(userID int64) (int64, error) {
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

func (uc *WordUC) GetRemindList(userID int64) ([]model.Word, error) {
	remindList, err := uc.WordRepo.GetRemindList(userID)
	if err != nil {
		return nil, err
	}
	return remindList, nil
}

func (uc *WordUC) FormatRemindList(words []model.Word) (string, error) {
	if len(words) == 0 {
		return "", apperrors.ErrNoWordsToRemind
	}

	var sb strings.Builder
	sb.WriteString("🌀 Слова на повторение:\n")

	for idx, word := range words {
		fmt.Fprintf(&sb, "%d. %s - %s\n", idx+1, word.Original, word.Translation)
	}
	return sb.String(), nil
}
