package bot

import (
	"langbrv/internal/infrastucture/transport/tgBot/handlers"

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
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextPage"),
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousPage"),
	),
)

var LastDictionaryPageKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousPage"),
	),
)

func ChooseDictionaryKeyboard(pageStatus handlers.DictionaryPageStatus) tgbotapi.InlineKeyboardMarkup {
	switch pageStatus {
	case handlers.FirstPage:
		return FirstDictionaryPageKeyboard
	case handlers.LastPage:
		return LastDictionaryPageKeyboard
	case handlers.SinglePage:
		return SingleDictionaryPageKeyboard
	}
	return MiddleDictionaryPageKeyboard
}
