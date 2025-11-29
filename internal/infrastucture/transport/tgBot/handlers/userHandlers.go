package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handlers) StartCommand(update tgbotapi.Update) string {
	// REMOVE: for testing
	//ADD: registration logic and greeting msg
	msg := "Привет, это LANGBRV, бот для изучения английского языка! Напиши /add_word чтобы добавить первое слово!"
	return msg
}
