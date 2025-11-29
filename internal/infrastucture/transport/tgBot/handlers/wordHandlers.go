package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handlers) AddWordCommand(update tgbotapi.Update) string {
	msg := "Enter word in format: word-translate"
	return msg
}
