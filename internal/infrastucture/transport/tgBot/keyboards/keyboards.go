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

func ChooseDictionaryKeyboard(pageStatus model.DictionaryPagePosition) interface{} {
	switch pageStatus {
	case model.FirstPage:
		return FirstDictionaryPageKeyboard
	case model.LastPage:
		return LastDictionaryPageKeyboard
	case model.SinglePage:
		return nil
	default:
		return MiddleDictionaryPageKeyboard
	}
}

func ChooseClosedRemindCardKeyboard(cardStatus model.RemindCardPosition) interface{} {
	switch cardStatus {
	case model.FirstCard:
		return FirstClosedRemindCardKeyboard
	case model.LastCard:
		return LastClosedRemindCardKeyboard
	case model.SingleCard:
		return nil
	default:
		return MiddleClosedRemindCardKeyboard
	}
}

func ChooseOpenedRemindCardKeyboard(cardStatus model.RemindCardPosition) interface{} {
	switch cardStatus {
	case model.FirstCard:
		return FirstOpenedRemindCardKeyboard
	case model.LastCard:
		return LastOpenedRemindCardKeyboard
	case model.SingleCard:
		return nil
	default:
		return MiddleOpenedRemindCardKeyboard
	}
}
