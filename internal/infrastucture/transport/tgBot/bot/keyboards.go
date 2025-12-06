package bot

import (
	"langbrv/internal/core/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var SingleDictionaryPageKeyboard = tgbotapi.NewInlineKeyboardMarkup()

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
		return SingleDictionaryPageKeyboard
	default:
		return MiddleDictionaryPageKeyboard
	}
}
