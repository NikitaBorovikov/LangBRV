package apperrors

import (
	"errors"
	"langbrv/internal/config"
)

var (
	ErrNoWordsInDictionary = errors.New("there are no words in the dictionary")
)

func HandleError(err error, msg *config.Errors) string {
	switch err {
	case ErrNoWordsInDictionary:
		return msg.NoWords
	default:
		return msg.Unknown
	}
}
