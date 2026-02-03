package bot

import (
	"errors"
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) GetDictionaryCommand(us *model.UserState, chatID int64) {
	us.RemindSession = nil
	us.DictionaryPage = model.NewDictionaryPage()

	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(us.UserID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNoWordsInDictionary) {
			b.handleNoWordsError(us, chatID)
			return
		}
		b.handleError(chatID, err)
		return
	}

	us.DictionaryPage.TotalPages = totalPages
	us.DictionaryPage.DeterminePosition()

	formattedPage, err := b.uc.DictionaryPageUC.FormatPage(us.UserID, us.DictionaryPage)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.ChooseDictionaryKeyboard(us.DictionaryPage.Position)
	us.DictionaryPage.MessageID = b.sendMessageWithKeyboard(chatID, formattedPage, keyboard)

	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetDictionaryCB(us *model.UserState, chatID int64) {
	us.RemindSession = nil
	us.DictionaryPage = model.NewDictionaryPage()

	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(us.UserID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNoWordsInDictionary) {
			b.handleNoWordsError(us, chatID)
			return
		}
		b.handleError(chatID, err)
		return
	}

	us.DictionaryPage.TotalPages = totalPages
	us.DictionaryPage.DeterminePosition()
	us.DictionaryPage.MessageID = us.LastMessageID

	if err := b.uc.UserStateUC.Save(us); err != nil {
		b.handleError(chatID, err)
		return
	}

	formattedPage, err := b.uc.DictionaryPageUC.FormatPage(us.UserID, us.DictionaryPage)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.ChooseDictionaryKeyboard(us.DictionaryPage.Position)
	b.updateMessage(chatID, us.LastMessageID, formattedPage, keyboard)
}

func (b *Bot) GetAnotherDictionaryPage(us *model.UserState, chatID int64, navigation model.Navigation) {
	us.RemindSession = nil
	us.DictionaryPage.ChangeCurrenctPage(navigation)

	if err := b.uc.UserStateUC.Save(us); err != nil {
		b.handleError(chatID, err)
		return
	}

	formattedPage, err := b.uc.DictionaryPageUC.FormatPage(us.UserID, us.DictionaryPage)
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	keyboard := keyboards.ChooseDictionaryKeyboard(us.DictionaryPage.Position)
	b.updateMessage(chatID, us.DictionaryPage.MessageID, formattedPage, keyboard)
}

func (b *Bot) handleNoWordsError(us *model.UserState, chatID int64) {
	us.DictionaryPage.MessageID = b.sendMessageWithKeyboard(chatID, b.msg.Errors.NoWords, keyboards.AddWordKeyboard)
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
	}
}
