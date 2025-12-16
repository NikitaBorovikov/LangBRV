package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) AddWordCommand(userID, chatID int64) {
	state := model.NewUserState(userID, model.AddWord, 0)

	if err := b.uc.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Info.AddWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) GetRemindListCommand(userID, chatID int64) {
	remindList, err := b.uc.WordUC.GetRemindList(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	remindMsg, err := b.uc.WordUC.FormatRemindList(remindList)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.sendMessage(chatID, remindMsg)
}

func (b *Bot) DeleteWordCommand(userID, chatID int64) {
	state := model.NewUserState(userID, model.DelWord, 0)

	if err := b.uc.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Info.DelWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) SaveWord(userID, chatID int64, text string) {
	userState, err := b.uc.UserStateUC.Get(userID)
	if err != nil || userState == nil {
		logrus.Error(err)
		msgText := b.msg.Errors.UnknownMsg
		b.sendMessage(chatID, msgText)
		return
	}

	req := dto.NewAddWordRequest(userID, text)
	word, err := req.ToDomainWord()
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	wordID, err := b.uc.WordUC.Add(word)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	msgText := b.msg.Success.WordAdded
	msgID := b.sendMessageWithKeyboard(chatID, msgText, keyboards.MainKeyboard)
	userState.LastMsgID = msgID
	logrus.Infof("word is saved: id = %s", wordID)
}

func (b *Bot) DeleteWord(userID, chatID int64, text string) {
	if err := b.uc.WordUC.Delete(userID, text); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Success.WordDeleted
	b.sendMessage(chatID, msgText)
}
