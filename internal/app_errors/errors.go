package apperrors

import (
	"errors"
	"langbrv/internal/config"
)

var (
	ErrNoWordsInDictionary = errors.New("there are no words in the dictionary")
	ErrNoWordsToRemind     = errors.New("there ate no words to remind")
	ErrWordNotFound        = errors.New("such word is not found")
)

func HandleError(err error, msg *config.Errors) string {
	switch err {
	case ErrNoWordsInDictionary:
		return msg.NoWords
	case ErrNoWordsToRemind:
		return msg.NoWordsToRemind
	case ErrWordNotFound:
		return msg.WordNotExists
	default:
		return msg.Unknown
	}
}
