package apperrors

import (
	"errors"
	"langbrv/internal/config"
)

var (
	ErrNoWordsInDictionary = errors.New("there are no words in the dictionary")
	ErrNoWordsToRemind     = errors.New("there are no words to remind")
	ErrWordNotFound        = errors.New("such word is not found")
	ErrUserStateNotFound   = errors.New("failed to find state with such userID")

	// Validation error
	ErrWordTooLong              = errors.New("validation error: word is too long")
	ErrIncorrectFormat          = errors.New("validation error: incorrect format")
	ErrIncorrectDeleteMsgFormat = errors.New("validation error: incorrect delete message format")
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
	case ErrIncorrectFormat:
		return msg.IncorrectFormat
	case ErrIncorrectDeleteMsgFormat:
		return msg.IncorrectDeleteMsgFormat
	default:
		return msg.Unknown
	}
}
