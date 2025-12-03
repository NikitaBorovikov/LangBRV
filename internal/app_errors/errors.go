package apperrors

import (
	"errors"
	"langbrv/internal/config"
)

var (
	ErrNoWordsInDictionary = errors.New("there are no words in the dictionary")
	ErrNoWordsToRemind     = errors.New("there ate no words to remind")
)

func HandleError(err error, msg *config.Errors) string {
	switch err {
	case ErrNoWordsInDictionary:
		return msg.NoWords
	case ErrNoWordsToRemind:
		return msg.NoWordsToRemind
	default:
		return msg.Unknown
	}
}
