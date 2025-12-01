package postgres

import (
	"errors"
	"langbrv/internal/core/model"

	"gorm.io/gorm"
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

func (r *WordRepo) GetAll(userID int64) ([]model.Word, error) {
	var words []model.Word
	result := r.db.Where("user_id = ?", userID).Order("last_seen DESC").Find(&words)
	return words, result.Error
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

func (r *WordRepo) Update(word *model.Word) error {
	result := r.db.Model(word).Where("id = ?", word.ID).Update("last_seen", word.LastSeen)
	return result.Error
}
