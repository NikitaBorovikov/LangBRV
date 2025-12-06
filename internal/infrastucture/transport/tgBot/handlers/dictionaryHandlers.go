package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) GetDictionaryCommand(update tgbotapi.Update) string {
	page := model.NewDictionaryPage(update.Message.From.ID)

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatDictionaryPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return formatedPage
}
