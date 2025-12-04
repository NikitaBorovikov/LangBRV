package apperrors

import (
	"errors"
	"langbrv/internal/config"
)

var (
	ErrNoWordsInDictionary = errors.New("there are no words in the dictionary")
	ErrNoWordsToRemind     = errors.New("there ate no words to remind")
	ErrWordNotFound        = errors.New("such word is not found")

	//validation error
	ErrWordTooLong      = errors.New("validation error: word is too long")
	ErrMissingSeparator = errors.New("validation error: missing separator in add-word request")
)

func HandleError(err error, msg *config.Errors) string {
	switch err {
	case ErrNoWordsInDictionary:
		return msg.NoWords
	case ErrNoWordsToRemind:
		return msg.NoWordsToRemind
	case ErrWordNotFound:
		return msg.WordNotExists
	case ErrWordTooLong:
		return msg.WordTooLong
	case ErrMissingSeparator:
		return msg.MissingSeparator
	default:
		return msg.Unknown
	}
}
