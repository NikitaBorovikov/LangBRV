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

// ключ - текущий уровень. значение - через сколько дней следующее повторение
var nextRepIn = map[uint8]uint8{
	1: 1,
	2: 2,
	3: 4,
	4: 6,
	5: 8,
	6: 10,
	7: 59,
}

type WordUC struct {
	WordRepo repository.WordRepo
}

func NewWordUC(wr repository.WordRepo) *WordUC {
	return &WordUC{
		WordRepo: wr,
	}
}

func (uc *WordUC) Add(word *model.Word) error {
	// Проверяем, есть ли уже такое слово в словаре
	existingWord, err := uc.WordRepo.FindByUserAndWord(word.UserID, word.Original)
	if err != nil {
		return err
	}

	// Если слово уже есть, то просто обновляем его поля
	if existingWord != nil {
		setDefaultWordFields(existingWord)
		err := uc.WordRepo.Update(existingWord)
		return err
	}

	// Если слова нет, то добавляем его
	setDefaultWordFields(word)
	word.CreatedAt = time.Now().UTC()

	if err := uc.WordRepo.Add(word); err != nil {
		return err
	}
	return nil
}

func (uc *WordUC) Update(word *model.Word, isRememberWell bool) error {
	if isRememberWell {
		word.NextRemind = time.Now().UTC().Add(time.Duration(nextRepIn[word.MemorizationLevel]) * 24 * time.Hour)
		word.MemorizationLevel += model.DefaultMemorizationLevelStep
	} else {
		word.NextRemind = time.Now().UTC().Add(time.Duration(nextRepIn[word.MemorizationLevel]) * 24 * time.Hour)
		word.MemorizationLevel = model.DefaultMemorizationLevel
	}

	if err := uc.WordRepo.Update(word); err != nil {
		return err
	}
	return nil
}

func (uc *WordUC) Delete(userID int64, word string) error {
	if err := dto.ValidateWord(word); err != nil {
		return err
	}

	word = strings.ToLower(word)
	if err := uc.WordRepo.Delete(userID, word); err != nil {
		return err
	}
	return nil
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
	sb.Grow(expectedPageSize)
	sb.WriteString("🌀 <b>Слова на повторение:</b>\n\n")

	for _, word := range words {
		fmt.Fprintf(&sb, "• %s - %s\n", word.Original, word.Translation)
	}
	return sb.String(), nil
}

func setDefaultWordFields(word *model.Word) {
	now := time.Now().UTC()
	word.LastSeen = now
	word.MemorizationLevel = model.DefaultMemorizationLevel
	word.NextRemind = now.Add(time.Duration(model.DefaultMemorizationLevel * 24 * time.Hour))
}
