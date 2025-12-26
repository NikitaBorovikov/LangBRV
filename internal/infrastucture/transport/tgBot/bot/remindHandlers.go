package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) GetRemindCardCommand(userID, chatID int64) {
	remindList, err := b.uc.WordUC.GetRemindList(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	card := model.NewRemindCard(userID, remindList)
	card.DeterminePosition()

	if err := b.uc.RemindCardUC.Save(card); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	keyboardType := keyboards.ChooseClosedRemindCardKeyboard(card.Position)

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*card)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	card.MessageID = b.sendMessageWithKeyboard(chatID, cardMsg, keyboardType)
}

func (b *Bot) GetAnotherRemindCard(userID, chatID int64, navigation Navigation) {
	card, err := b.uc.RemindCardUC.Get(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	if navigation == Next {
		card.CurrentCard++
	} else {
		card.CurrentCard--
	}

	card.DeterminePosition()
	keyboard := keyboards.ChooseClosedRemindCardKeyboard(card.Position)

	if err := b.uc.RemindCardUC.Save(card); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*card)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, card.MessageID, cardMsg, keyboard)
}

func (b *Bot) ShowRemindCard(userID, chatID int64) {
	card, err := b.uc.RemindCardUC.Get(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	card.DeterminePosition()
	keyboard := keyboards.ChooseOpenedRemindCardKeyboard(card.Position)

	cardMsg, err := b.uc.RemindCardUC.FormatOpenedRemindCard(*card)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, card.MessageID, cardMsg, keyboard)
}
