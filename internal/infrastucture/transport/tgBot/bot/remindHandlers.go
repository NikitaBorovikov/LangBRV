package bot

import (
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"
	"time"

	"github.com/sirupsen/logrus"
)

func (b *Bot) StartRemindSession(us *model.UserState, chatID int64) {
	remindList, err := b.uc.WordUC.GetRemindList(us.UserID)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	us.RemindSession = model.NewRemindSession(remindList)

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.ClosedRemindCardKeyboard
	us.RemindSession.MessageID = b.sendMessageWithKeyboard(chatID, cardMsg, keyboard)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetNextRemindCard(us *model.UserState, chatID int64, isRememberWell bool) {
	if err := b.updatePreviousCardInfo(us, isRememberWell); err != nil {
		logrus.Errorf("failed to update word: %v", err)
	}

	// If the previous card was the last or only one, we display a message indicating the end of the remind session
	if us.RemindSession.Position == model.Last || us.RemindSession.Position == model.Single {
		b.handleLastOrSingleRemindCard(us, chatID)
		return
	}

	us.RemindSession.GoToNextCard()
	us.DictionaryPage = nil

	if err := b.uc.UserStateUC.Save(us); err != nil {
		b.handleError(chatID, err)
		return
	}

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		b.handleError(chatID, err)
		return
	}
	keyboard := keyboards.ClosedRemindCardKeyboard
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
}

func (b *Bot) ShowRemindCard(us *model.UserState, chatID int64) {
	us.RemindSession.DeterminePosition()

	cardMsg, err := b.uc.RemindCardUC.FormatOpenedRemindCard(*us.RemindSession)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.OpenedRemindCardKeyboard
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
}

func (b *Bot) RepeatRemindSession(us *model.UserState, chatID int64) {
	us.RemindSession.CurrentCard = model.DefaultRemindCardNumber
	us.RemindSession.DeterminePosition()

	cardMsg, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.ClosedRemindCardKeyboard
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)

	us.DictionaryPage = nil
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) updatePreviousCardInfo(us *model.UserState, isRememberWell bool) error {
	previousCardIdx := us.RemindSession.CurrentCard - 1
	previousWord := &us.RemindSession.Words[previousCardIdx]

	// Prevent the word data from being updated during retraining.
	lastSeen := previousWord.LastSeen.Truncate(24 * time.Hour)
	today := time.Now().UTC().Truncate(24 * time.Hour)
	if !lastSeen.Equal(today) {
		if err := b.uc.WordUC.Update(previousWord, isRememberWell); err != nil {
			return err
		}
	}
	return nil
}

func (b *Bot) handleLastOrSingleRemindCard(us *model.UserState, chatID int64) {
	keyboard := keyboards.RemindSessionIsOverKeyboard
	cardMsg := b.msg.Info.RemindSessionIsOver
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
}
