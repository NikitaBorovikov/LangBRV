package keyboards

import (
	"langbrv/internal/core/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var StartKeyboard = tgbotapi.NewInlineKeyboardMarkup(
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

var ClosedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать перевод", "showWord"),
	),
)

var SingleClosedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать перевод", "showWord"),
	),
)

var OpenedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("👎", "rememberBadly"),
		tgbotapi.NewInlineKeyboardButtonData("👍", "rememberWell"),
	),
)

var RemindSessionIsOverKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать список слов", "showRemindList"),
	),
)

func ChooseDictionaryKeyboard(pageStatus model.Position) interface{} {
	switch pageStatus {
	case model.First:
		return FirstDictionaryPageKeyboard
	case model.Last:
		return LastDictionaryPageKeyboard
	case model.Single:
		return nil
	default:
		return MiddleDictionaryPageKeyboard
	}
}
