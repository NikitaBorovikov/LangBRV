package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) GetDictionaryCommand(us *model.UserState, chatID int64) {
	page := model.NewDictionaryPage()
	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(us.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		if errMsgText == b.msg.Errors.NoWords {
			page.MessageID = b.sendMessageWithKeyboard(chatID, errMsgText, keyboards.AddWordKeyboard)
			return
		}
		b.sendMessage(chatID, errMsgText)
		return
	}

	page.TotalPages = totalPages
	page.DeterminePosition()

	keyboard := keyboards.ChooseDictionaryKeyboard(page.Position)

	us.DictionaryPage = page
	us.RemindCard = nil
	us.Mode = model.ViewDictionaryMode

	formatedPage, err := b.uc.DictionaryPageUC.FormatPage(us.UserID, page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	us.DictionaryPage.MessageID = b.sendMessageWithKeyboard(chatID, formatedPage, keyboard)

	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		return
	}
}

func (b *Bot) GetDictionaryCB(us *model.UserState, chatID int64) {
	page := model.NewDictionaryPage()
	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(us.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		if errMsgText == b.msg.Errors.NoWords {
			us.DictionaryPage.MessageID = b.sendMessageWithKeyboard(chatID, errMsgText, keyboards.AddWordKeyboard)
			return
		}
		b.sendMessage(chatID, errMsgText)
		return
	}

	page.TotalPages = totalPages
	page.DeterminePosition()
	page.MessageID = us.LastMessageID

	keyboard := keyboards.ChooseDictionaryKeyboard(page.Position)

	us.DictionaryPage = page
	us.RemindCard = nil
	us.Mode = model.ViewDictionaryMode

	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	formatedPage, err := b.uc.DictionaryPageUC.FormatPage(us.UserID, page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.LastMessageID, formatedPage, keyboard)
}

func (b *Bot) GetAnotherDictionaryPage(us *model.UserState, chatID int64, navigation Navigation) {
	if navigation == Next {
		us.DictionaryPage.CurrentPage++
	} else {
		us.DictionaryPage.CurrentPage--
	}

	us.DictionaryPage.DeterminePosition()
	keyboard := keyboards.ChooseDictionaryKeyboard(us.DictionaryPage.Position)

	us.RemindCard = nil
	us.Mode = model.ViewDictionaryMode
	if err := b.uc.UserStateUC.Save(us); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	formatedPage, err := b.uc.DictionaryPageUC.FormatPage(us.UserID, us.DictionaryPage)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, us.DictionaryPage.MessageID, formatedPage, keyboard)
}
