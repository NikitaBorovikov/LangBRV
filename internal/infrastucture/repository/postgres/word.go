package postgres

import (
	"database/sql"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"

	"github.com/jmoiron/sqlx"
)

type WordRepo struct {
	db *sqlx.DB
}

func NewWordRepo(db *sqlx.DB) *WordRepo {
	return &WordRepo{
		db: db,
	}
}

func (r *WordRepo) Add(word *model.Word) (string, error) {
	query := `INSERT INTO words (user_id, original, translation, last_seen, created_at)
	VALUES (:user_id, :original, :translation, :last_seen, :created_at)
	RETURNING id`

	rows, err := r.db.NamedQuery(query, word)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&word.ID); err != nil {
			return "", err
		}
	}
	return word.ID, err
}

func (r *WordRepo) GetDictionaryWordsByPage(userID, pageNum, wordsPerPage int64) ([]model.Word, error) {
	var words []model.Word
	offset := (pageNum - 1) * wordsPerPage

	query := `SELECT original, translation FROM words WHERE user_id = $1
	ORDER BY last_seen DESC
	LIMIT $2 OFFSET $3`

	if err := r.db.Select(&words, query, userID, wordsPerPage, offset); err != nil {
		return nil, err
	}
	return words, nil
}

func (r *WordRepo) GetAmountOfWords(userID int64) (int64, error) {
	var wordsAmount int64

	query := `SELECT COUNT(*) FROM words WHERE user_id = $1`

	if err := r.db.Get(&wordsAmount, query, userID); err != nil {
		return 0, err
	}
	return wordsAmount, nil
}

func (r *WordRepo) FindByUserAndWord(userID int64, word string) (*model.Word, error) {
	var existingWord model.Word

	query := `SELECT * FROM words WHERE user_id = $1 AND original = $2 LIMIT 1`

	err := r.db.Get(&existingWord, query, userID, word)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Если не нашли слово - это не ошибка
		}
		return nil, err
	}
	return &existingWord, nil
}

func (r *WordRepo) GetRemindList(userID int64) ([]model.Word, error) {
	var remindWords []model.Word

	query := `SELECT original, translation FROM words 
	WHERE user_id = $1 AND CURRENT_DATE - last_seen::date IN (1, 3, 10, 30, 90)
	ORDER BY last_seen ASC`

	err := r.db.Select(&remindWords, query, userID)
	if err != nil {
		return nil, err
	}
	return remindWords, nil
}

func (r *WordRepo) Update(word *model.Word) error {
	query := `UPDATE words SET last_seen = $1 WHERE id = $2`
	_, err := r.db.Exec(query, word.LastSeen, word.ID)
	return err
}

func (r *WordRepo) Delete(userID int64, word string) error {
	query := `DELETE FROM words WHERE user_id = $1 AND (original = $2 OR translation = $2)`

	result, err := r.db.Exec(query, userID, word)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrWordNotFound
	}
	return nil
}
