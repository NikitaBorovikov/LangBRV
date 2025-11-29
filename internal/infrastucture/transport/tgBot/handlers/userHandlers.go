package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handlers) StartCommand(update tgbotapi.Update) string {
	// REMOVE: for testing
	//ADD: registration logic and greeting msg
	msg := "Hello, this is LANGBRV bot: bot for learning English"
	return msg
}
