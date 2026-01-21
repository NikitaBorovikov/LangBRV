package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"
	"time"

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

	us.RemindSession = model.NewRemindSession(remindList)
	us.RemindSession.DeterminePosition()

	keyboard := keyboards.ClosedRemindCardKeyboard

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	us.RemindSession.MessageID = b.sendMessageWithKeyboard(chatID, cardMsg, keyboard)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetNextRemindCard(us *model.UserState, chatID int64, isRememberWell bool) {
	// меняем memorizationLevel и newRemind для предыдущей карточки
	previousCardIdx := us.RemindSession.CurrentCard - 1
	word := us.RemindSession.Words[previousCardIdx]

	// предотварщаем обновление данных о слове при повторной тренировке
	lastSeenDay := us.RemindSession.Words[previousCardIdx].LastSeen.Day()
	today := time.Now().UTC().Day()
	if lastSeenDay != today {
		if err := b.uc.WordUC.Update(&word, isRememberWell); err != nil {
			logrus.Errorf("failed to update word: %v", err)
		}
	}

	// Если предыдущая карточка была последней или единственной - показываем сообщение о завершении тренировки
	if us.RemindSession.Position == model.Last || us.RemindSession.Position == model.Single {
		keyboard := keyboards.RemindSessionIsOverKeyboard
		cardMsg := b.msg.Info.RemindSessionIsOver
		b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
		return
	}

	// Переходим к следущей карточке
	us.RemindSession.CurrentCard++
	us.RemindSession.DeterminePosition()
	keyboard := keyboards.ClosedRemindCardKeyboard

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
}

func (b *Bot) ShowRemindCard(us *model.UserState, chatID int64) {
	us.RemindSession.DeterminePosition()
	keyboard := keyboards.OpenedRemindCardKeyboard

	cardMsg, err := b.uc.RemindCardUC.FormatOpenedRemindCard(*us.RemindSession)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
}

func (b *Bot) RepeatRemindSession(us *model.UserState, chatID int64) {
	us.RemindSession.CurrentCard = model.DefaultRemindCardNumber
	keyboard := keyboards.ClosedRemindCardKeyboard

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}
