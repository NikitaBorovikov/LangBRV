package postgres

import (
	"errors"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"

	"gorm.io/gorm"
)

// REMOVE: for testing
const (
	wordsPerPage = 5
)

type WordRepo struct {
	db *gorm.DB
}

func NewWordRepo(db *gorm.DB) *WordRepo {
	return &WordRepo{
		db: db,
	}
}

func (r *WordRepo) Add(word *model.Word) (string, error) {
	result := r.db.Create(word)
	return word.ID, result.Error
}

func (r *WordRepo) GetDictionaryWordsByPage(userID, pageNum int64) ([]model.Word, error) {
	var words []model.Word
	offset := (pageNum - 1) * wordsPerPage

	err := r.db.Where("user_id = ?", userID).Order("last_seen DESC").Offset(int(offset)).Limit(wordsPerPage).Find(&words).Error
	if err != nil {
		return nil, err
	}

	return words, nil
}

func (r *WordRepo) GetAmountOfWords(userID int64) (int64, error) {
	var wordsAmount int64

	err := r.db.Model(&model.Word{}).Where("user_id = ?", userID).Count(&wordsAmount).Error
	if err != nil {
		return 0, err
	}
	return wordsAmount, nil
}

func (r *WordRepo) FindByUserAndWord(userID int64, word string) (*model.Word, error) {
	var existingWord model.Word

	err := r.db.Where("user_id = ? AND original = ?", userID, word).First(&existingWord).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Если не нашли слово - это не ошибка
		}
		return nil, err
	}
	return &existingWord, nil
}

func (r *WordRepo) GetRemindList(userID int64) ([]model.Word, error) {
	var remindWords []model.Word
	remindIntervals := []int{1, 3, 10, 30, 90}

	err := r.db.Where("user_id = ? AND ((CURRENT_DATE - last_seen::date) IN ?)", userID, remindIntervals).Order("last_seen ASC").Find(&remindWords).Error
	if err != nil {
		return nil, err
	}
	return remindWords, nil
}

func (r *WordRepo) Update(word *model.Word) error {
	result := r.db.Model(word).Where("id = ?", word.ID).Update("last_seen", word.LastSeen)
	return result.Error
}

func (r *WordRepo) Delete(userID int64, word string) error {
	result := r.db.Where("user_id = ? AND original = ?", userID, word).Or("user_id = ? AND translation = ?", userID, word).Delete(&model.Word{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrWordNotFound
	}
	return nil
}
