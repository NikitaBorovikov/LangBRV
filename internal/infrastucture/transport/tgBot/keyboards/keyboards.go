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

var FirstClosedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать перевод", "showWord"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextCard"),
	),
)

var MiddleClosedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать перевод", "showWord"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousCard"),
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextCard"),
	),
)

var LastClosedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Показать перевод", "showWord"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousCard"),
	),
)

var FirstOpenedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextCard"),
	),
)

var MiddleOpenedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousCard"),
		tgbotapi.NewInlineKeyboardButtonData("➡️", "nextCard"),
	),
)

var LastOpenedRemindCardKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("⬅️", "previousCard"),
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

func ChooseClosedRemindCardKeyboard(cardStatus model.Position) interface{} {
	switch cardStatus {
	case model.First:
		return FirstClosedRemindCardKeyboard
	case model.Last:
		return LastClosedRemindCardKeyboard
	case model.Single:
		return nil
	default:
		return MiddleClosedRemindCardKeyboard
	}
}

func ChooseOpenedRemindCardKeyboard(cardStatus model.Position) interface{} {
	switch cardStatus {
	case model.First:
		return FirstOpenedRemindCardKeyboard
	case model.Last:
		return LastOpenedRemindCardKeyboard
	case model.Single:
		return nil
	default:
		return MiddleOpenedRemindCardKeyboard
	}
}
