package bot

import (
	"langbrv/internal/core/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AddFirstWordKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить первое слово", "addWord"),
	),
)

var AddWordKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить слово", "addWord"),
	),
)

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Добавить ещё", "addWord"),
		tgbotapi.NewInlineKeyboardButtonData("Мой словарь", "getDictionary"),
	),
)

var FirstDictionaryPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextPage"),
	),
)

var MiddleDictionaryPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousPage"),
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextPage"),
	),
)

var LastDictionaryPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousPage"),
	),
)

func ChooseDictionaryKeyboard(pageStatus model.DictionaryPageStatus) tgbotapi.InlineKeyboardMarkup {
	switch pageStatus {
	case model.FirstPage:
		return FirstDictionaryPageKeyboard
	case model.LastPage:
		return LastDictionaryPageKeyboard
	case model.SinglePage:
		return AddWordKeyboard
	default:
		return MiddleDictionaryPageKeyboard
	}
}
