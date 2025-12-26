package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) GetRemindCardCommand(us *model.UserState, chatID int64) {
	remindList, err := b.uc.WordUC.GetRemindList(us.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	us.RemindCard = model.NewRemindCard(remindList)
	us.RemindCard.DeterminePosition()
	us.Mode = model.RemidMode

	keyboardType := keyboards.ChooseClosedRemindCardKeyboard(us.RemindCard.Position)

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindCard)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	us.RemindCard.MessageID = b.sendMessageWithKeyboard(chatID, cardMsg, keyboardType)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetAnotherRemindCard(us *model.UserState, chatID int64, navigation Navigation) {
	if navigation == Next {
		us.RemindCard.CurrentCard++
	} else {
		us.RemindCard.CurrentCard--
	}

	us.RemindCard.DeterminePosition()
	us.Mode = model.RemidMode
	keyboard := keyboards.ChooseClosedRemindCardKeyboard(us.RemindCard.Position)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindCard)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.RemindCard.MessageID, cardMsg, keyboard)
}

func (b *Bot) ShowRemindCard(us *model.UserState, chatID int64) {
	us.RemindCard.DeterminePosition()
	us.Mode = model.RemidMode
	keyboard := keyboards.ChooseOpenedRemindCardKeyboard(us.RemindCard.Position)

	cardMsg, err := b.uc.RemindCardUC.FormatOpenedRemindCard(*us.RemindCard)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.RemindCard.MessageID, cardMsg, keyboard)
}
