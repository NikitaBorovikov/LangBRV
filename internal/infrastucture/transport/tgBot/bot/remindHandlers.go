package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *Bot) GetRemindCardCommand(userID, chatID int64) {
	card := model.NewRemindCard(userID)
	remindList, err := b.uc.WordUC.GetRemindList(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	card.TotalCards = len(remindList) // Одно слово на карточке
	card.Words = remindList
	card.DetermineStatus()

	if err := b.uc.RemindCardUC.Save(card); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	keyboardType := keyboards.ChooseRemindCardKeyboard(card.Status)

	cardMsg, err := b.uc.RemindCardUC.FormatRemindCard(*card)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	card.RemindMsgID = b.sendMessageWithKeyboard(chatID, cardMsg, keyboardType)
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

	card.DetermineStatus()
	keyboardType := keyboards.ChooseRemindCardKeyboard(card.Status)
	keyboard, ok := keyboardType.(tgbotapi.InlineKeyboardMarkup)
	if !ok {
		b.sendMessage(chatID, b.msg.Errors.Unknown)
		return
	}

	if err := b.uc.RemindCardUC.Save(card); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	cardMsg, err := b.uc.RemindCardUC.FormatRemindCard(*card)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, card.RemindMsgID, cardMsg, keyboard)
}
