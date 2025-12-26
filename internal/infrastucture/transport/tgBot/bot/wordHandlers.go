package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) AddWord(us *model.UserState, chatID int64) {
	us.Mode = model.AddMode
	msgText := b.msg.Info.AddWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) DeleteWordCommand(us *model.UserState, chatID int64) {
	us.Mode = model.DeleteMode
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Info.DelWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) SaveWord(us *model.UserState, chatID int64, text string) {
	req := dto.NewAddWordRequest(us.UserID, text)
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
	us.LastMessageID = msgID
}

func (b *Bot) DeleteWord(us *model.UserState, chatID int64, text string) {
	defer func() {
		us.Mode = model.AddMode // Выключаем режим удаления
	}()

	if err := b.uc.WordUC.Delete(us.UserID, text); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Success.WordDeleted
	msgID := b.sendMessageWithKeyboard(chatID, msgText, keyboards.MainKeyboard)
	us.LastMessageID = msgID
}
