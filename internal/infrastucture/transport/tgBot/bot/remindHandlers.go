package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) StartRemindSession(us *model.UserState, chatID int64) {
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

func (b *Bot) GetNextRemindCard(us *model.UserState, chatID int64, isRememberWell bool) {
	// меняем memorizationLevel и newRemind для предыдущей карточки
	word := us.RemindCard.Words[us.RemindCard.CurrentCard-1]
	if err := b.uc.WordUC.Update(&word, isRememberWell); err != nil {
		logrus.Errorf("failed to update word: %v", err)
	}

	// Если предыдущая карточка была последней или единственной - показываем сообщение о завершении тренировки
	if us.RemindCard.Position == model.Last || us.RemindCard.Position == model.Single {
		keyboard := keyboards.RemindSessionIsOverKeyboard
		cardMsg := b.msg.Info.RemindSessionIsOver
		b.updateMessage(chatID, us.RemindCard.MessageID, cardMsg, keyboard)
		return
	}

	// Переходим к следущей карточке
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

func (b *Bot) RepeatRemindSession(us *model.UserState, chatID int64) {
	us.Mode = model.RemidMode
	us.RemindCard.CurrentCard = model.DefaultCardNumber
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
