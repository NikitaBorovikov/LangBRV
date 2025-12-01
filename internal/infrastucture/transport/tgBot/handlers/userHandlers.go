package handlers

import (
	"langbrv/internal/infrastucture/transport/tgBot/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) StartCommand(update tgbotapi.Update) string {
	req := dto.NewRegistrationRequest(update.Message.From.ID, update.Message.From.UserName)
	user := req.ToDomainUser()

	if err := h.UseCases.UserUC.CreateOrUpdate(user); err != nil {
		return "Ошибка сохранения. Попробуй позже!"
	}

	msg := "Привет, это LANGBRV, бот для изучения английского языка! Напиши /add_word чтобы добавить первое слово!"
	return msg
}
