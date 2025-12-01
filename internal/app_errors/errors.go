package apperrors

import "errors"

var (
	ErrNoWordsInDictionary = errors.New("there are no words in the dictionary")
)

// TODO: поместить тексты ошибок в конфиг
const (
	ErrNoWordsInDictionaryMsg = "В словаре пока нет слов.\nЧтобы добавить слово напиши /add_word"
	UnknowErrMsg              = "Неизвестная ошибка"
)

func HandleError(err error) string {
	switch err {
	case ErrNoWordsInDictionary:
		return ErrNoWordsInDictionaryMsg
	default:
		return UnknowErrMsg
	}
}
