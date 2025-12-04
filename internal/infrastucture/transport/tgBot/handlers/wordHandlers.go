package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) AddWordCommand(update tgbotapi.Update) string {
	state := model.NewUserState(update.Message.Chat.ID, model.AddWord)

	if err := h.UseCases.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Info.AddWord
}

func (h *Handlers) GetDictionaryCommand(update tgbotapi.Update) string {
	words, err := h.UseCases.WordUC.GetAll(update.Message.From.ID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}

	dictionary, err := h.UseCases.WordUC.FormatDictionary(words)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return dictionary
}

func (h *Handlers) GetRemindListCommand(update tgbotapi.Update) string {
	remindList, err := h.UseCases.WordUC.GetRemindList(update.Message.From.ID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}

	remindMsg, err := h.UseCases.WordUC.FormatRemindList(remindList)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return remindMsg
}

func (h *Handlers) DeleteWordCommand(update tgbotapi.Update) string {
	state := model.NewUserState(update.Message.From.ID, model.DelWord)

	if err := h.UseCases.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Info.DelWord
}

func (h *Handlers) SaveWord(update tgbotapi.Update) string {
	req := dto.NewAddWordRequest(update.Message.From.ID, update.Message.Text)
	word, err := req.ToDomainWord()
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}

	wordID, err := h.UseCases.WordUC.Add(word)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	logrus.Infof("word is saved: id = %s", wordID)
	return h.Msg.Success.WordAdded
}
