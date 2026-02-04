package usecases

import (
	"fmt"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/core/repository"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	DeleteWordsSeparator = ","
)

// key - current level. value - how many days until the next repetition
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
	// Checking if this word is already in the dictionary.
	existingWord, err := uc.WordRepo.FindByUserAndWord(word.UserID, word.Original)
	if err != nil {
		return err
	}

	setDefaultWordFields(word)

	// If the word already exists, we simply update its fields.
	if existingWord != nil {
		err := uc.WordRepo.Update(existingWord)
		return err
	}

	// If the word is not present, we add it
	word.CreatedAt = time.Now().UTC()

	if err := uc.WordRepo.Add(word); err != nil {
		return err
	}
	return nil
}

func (uc *WordUC) Update(word *model.Word, isRememberWell bool) error {
	now := time.Now().UTC()

	if isRememberWell {
		word.NextRemind = now.Add(time.Duration(nextRepIn[word.MemorizationLevel]) * 24 * time.Hour)
		word.MemorizationLevel += model.DefaultMemorizationLevelStep
	} else {
		word.NextRemind = now.Add(time.Duration(nextRepIn[word.MemorizationLevel]) * 24 * time.Hour)
		word.MemorizationLevel = model.DefaultMemorizationLevel
	}

	word.LastSeen = now

	if err := uc.WordRepo.Update(word); err != nil {
		return err
	}
	return nil
}

func (uc *WordUC) Delete(userID int64, words string) (int, error) {
	wordsArr := strings.Split(words, DeleteWordsSeparator)

	if len(wordsArr) == 0 || len(wordsArr) >= 25 {
		return 0, apperrors.ErrIncorrectDeleteMsgFormat
	}

	var successfulDeletionsCounter int
	for _, word := range wordsArr {
		if err := dto.ValidateWord(word); err != nil {
			logrus.Errorf("validation word error: %v", err)
			continue
		}

		word = strings.TrimSpace(strings.ToLower(word))

		if err := uc.WordRepo.Delete(userID, word); err != nil {
			logrus.Errorf("failed to delete word: %v", err)
			continue
		}
		successfulDeletionsCounter++
	}
	return successfulDeletionsCounter, nil
}

func (uc *WordUC) GetRemindList(userID int64) ([]model.Word, error) {
	remindList, err := uc.WordRepo.GetRemindList(userID)
	return remindList, err
}

func (uc *WordUC) GetListOfRemindedWords(userID int64) ([]model.Word, error) {
	remindList, err := uc.WordRepo.GetListOfRemindedWords(userID)
	return remindList, err
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
