package dto

import apperrors "langbrv/internal/app_errors"

const (
	maxWordLength = 255
)

func ValidateWord(word string) error {
	if len(word) >= maxWordLength {
		return apperrors.ErrWordTooLong
	}
	return nil
}
