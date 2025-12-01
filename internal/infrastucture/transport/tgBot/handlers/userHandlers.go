package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/infrastucture/transport/tgBot/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) StartCommand(update tgbotapi.Update) string {
	req := dto.NewRegistrationRequest(update.Message.From.ID, update.Message.From.UserName)
	user := req.ToDomainUser()

	if err := h.UseCases.UserUC.CreateOrUpdate(user); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Info.Start
}
