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

	keyboard := keyboards.ClosedRemindCardKeyboard

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindCard)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	us.RemindCard.MessageID = b.sendMessageWithKeyboard(chatID, cardMsg, keyboard)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetAnotherRemindCard(us *model.UserState, chatID int64, rememberStatus RememerStatus) {
	// Если предыдущая карточка была последней - показываем сообщение о завершении тренировки
	if us.RemindCard.Position == model.Last {
		keyboard := keyboards.RemindSessionIsOverKeyboard
		cardMsg := b.msg.Info.RemindSessionIsOver
		b.updateMessage(chatID, us.RemindCard.MessageID, cardMsg, keyboard)
		return
	}

	// change memorizationLevel

	logrus.Info(us.RemindCard)

	us.RemindCard.CurrentCard++
	us.RemindCard.DeterminePosition()
	us.Mode = model.RemidMode
	keyboard := keyboards.ClosedRemindCardKeyboard

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
	keyboard := keyboards.OpenedRemindCardKeyboard

	cardMsg, err := b.uc.RemindCardUC.FormatOpenedRemindCard(*us.RemindCard)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.RemindCard.MessageID, cardMsg, keyboard)
}

func (b *Bot) GetAnotherRemindSession(us *model.UserState, chatID int64) {
	us.Mode = model.RemidMode
	us.RemindCard.CurrentCard = 1
	keyboard := keyboards.ClosedRemindCardKeyboard

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindCard)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	b.updateMessage(chatID, us.RemindCard.MessageID, cardMsg, keyboard)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}
