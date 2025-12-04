package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"strings"
	"time"
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

func (uc *WordUC) GetAll(userID int64) ([]model.Word, error) {
	words, err := uc.WordRepo.GetAll(userID)
	if err != nil {
		return nil, err
	}
	return words, err
}

func (uc *WordUC) DeleteWord(userID int64, word string) error {
	if len(word) >= 255 {
		return apperrors.ErrNoWordsInDictionary // Replace error
	}

	if err := uc.WordRepo.DeleteWord(userID, word); err != nil {
		return err
	}
	return nil
}

func (uc *WordUC) FormatDictionary(words []model.Word) (string, error) {
	if len(words) == 0 {
		return "", apperrors.ErrNoWordsInDictionary
	}

	//TODO: добавить предворительное выделение памяти
	var sb strings.Builder
	sb.WriteString("Твой словарь:\n")

	for idx, word := range words {
		fmt.Fprintf(&sb, "%d. %s - %s\n", idx+1, word.Original, word.Translation)
	}
	return sb.String(), nil
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
	sb.WriteString("Слова на повторение:\n")

	for idx, word := range words {
		fmt.Fprintf(&sb, "%d. %s - %s\n", idx+1, word.Original, word.Translation)
	}
	return sb.String(), nil
}
