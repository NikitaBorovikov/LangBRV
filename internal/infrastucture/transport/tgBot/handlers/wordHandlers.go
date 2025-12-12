package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"

	"github.com/sirupsen/logrus"
)

func (h *Handlers) AddWordCommand(userID int64) string {
	state := model.NewUserState(userID, model.AddWord)

	if err := h.UseCases.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Info.AddWord
}

func (h *Handlers) GetRemindListCommand(userID int64) string {
	remindList, err := h.UseCases.WordUC.GetRemindList(userID)
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

func (h *Handlers) DeleteWordCommand(userID int64) string {
	state := model.NewUserState(userID, model.DelWord)

	if err := h.UseCases.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Info.DelWord
}

func (h *Handlers) SaveWord(userID int64, msgText string) string {
	req := dto.NewAddWordRequest(userID, msgText)
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

func (h *Handlers) DeleteWord(userID int64, msgText string) string {
	if err := h.UseCases.WordUC.Delete(userID, msgText); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Success.WordDeleted
}
