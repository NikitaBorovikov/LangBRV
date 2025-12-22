package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) AddWord(userID, chatID int64) {
	state := model.NewUserState(userID, false, 0)

	if err := b.uc.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Info.AddWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) DeleteWordCommand(userID, chatID int64) {
	state := model.NewUserState(userID, true, 0)

	if err := b.uc.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Info.DelWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) SaveWord(state *model.UserState, chatID int64, text string) {
	req := dto.NewAddWordRequest(state.UserID, text)
	word, err := req.ToDomainWord()
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	if err := b.uc.WordUC.Add(word); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	msgText := b.msg.Success.WordAdded
	msgID := b.sendMessageWithKeyboard(chatID, msgText, keyboards.MainKeyboard)
	state.LastMsgID = msgID
}

func (b *Bot) DeleteWord(state *model.UserState, chatID int64, text string) {
	defer func() {
		state.DeleteMode = false // Выключаем режим удаления
	}()

	if err := b.uc.WordUC.Delete(state.UserID, text); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Success.WordDeleted
	msgID := b.sendMessageWithKeyboard(chatID, msgText, keyboards.MainKeyboard)
	state.LastMsgID = msgID
}
