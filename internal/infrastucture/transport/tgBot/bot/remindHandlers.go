package bot

import (
	"errors"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) StartRemindSession(us *model.UserState, chatID int64) {
	us.DictionaryPage = nil

	remindList, err := b.uc.WordUC.GetRemindList(us.UserID)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	us.RemindSession = model.NewRemindSession(remindList)

	formattedCard, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if errors.Is(err, apperrors.ErrNoWordsToRemind) {
		b.handleEmptyRemindList(us, chatID)
		return
	}
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.ClosedRemindCardKeyboard
	us.RemindSession.MessageID = b.sendMessageWithKeyboard(chatID, formattedCard, keyboard)

	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetNextRemindCard(us *model.UserState, chatID int64, isRememberWell bool) {
	us.DictionaryPage = nil

	if err := b.updatePreviousCardInfo(us, isRememberWell); err != nil {
		logrus.Errorf("failed to update word: %v", err)
	}

	// If the previous card was the last or only one, we display a message indicating the end of the remind session
	if us.RemindSession.Position == model.Last || us.RemindSession.Position == model.Single {
		b.handleLastOrSingleRemindCard(us, chatID)
		return
	}

	us.RemindSession.GoToNextCard()

	if err := b.uc.UserStateUC.Save(us); err != nil {
		b.handleError(chatID, err)
		return
	}

	formattedCard, err := b.uc.RemindCardUC.FormatClosedRemindCard(*us.RemindSession)
	if err != nil {
		b.handleError(chatID, err)
		return
	}
	keyboard := keyboards.ClosedRemindCardKeyboard
	b.updateMessage(chatID, us.RemindSession.MessageID, formattedCard, keyboard)
}

func (b *Bot) ShowRemindCard(us *model.UserState, chatID int64) {
	us.RemindSession.DeterminePosition()

	formattedCard, err := b.uc.RemindCardUC.FormatOpenedRemindCard(*us.RemindSession)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.OpenedRemindCardKeyboard
	b.updateMessage(chatID, us.RemindSession.MessageID, formattedCard, keyboard)
}

func (b *Bot) ShowListOfRemindedWords(us *model.UserState, chatID int64) {
	us.DictionaryPage = nil

	remindList, err := b.uc.WordUC.GetListOfRemindedWords(us.UserID)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	formattedList, err := b.uc.RemindCardUC.FormatListOfRemindedWords(remindList)

	b.updateMessage(chatID, us.RemindSession.MessageID, formattedList, nil)
}

func (b *Bot) updatePreviousCardInfo(us *model.UserState, isRememberWell bool) error {
	previousCardIdx := us.RemindSession.CurrentCard - 1
	previousWord := &us.RemindSession.Words[previousCardIdx]

	if err := b.uc.WordUC.Update(previousWord, isRememberWell); err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleLastOrSingleRemindCard(us *model.UserState, chatID int64) {
	keyboard := keyboards.RemindSessionIsOverKeyboard
	cardMsg := b.msg.Info.RemindSessionIsOver
	b.updateMessage(chatID, us.RemindSession.MessageID, cardMsg, keyboard)
}

func (b *Bot) handleEmptyRemindList(us *model.UserState, chatID int64) {
	listOfRemindedWords, err := b.uc.WordUC.GetListOfRemindedWords(us.UserID)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	if len(listOfRemindedWords) == 0 {
		msgText := b.msg.Errors.NoWordsToRemind
		b.sendMessage(chatID, msgText)
		return
	}

	formattedList, err := b.uc.RemindCardUC.FormatListOfRemindedWords(listOfRemindedWords)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	b.sendMessage(chatID, formattedList)
}
