package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *Bot) GetDictionaryCommand(userID, chatID int64) {
	page := model.NewDictionaryPage(userID)
	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		if errMsgText == b.msg.Errors.NoWords {
			page.DictionaryMsgID = b.sendMessageWithKeyboard(chatID, errMsgText, keyboards.AddWordKeyboard)
			return
		}
		b.sendMessage(chatID, errMsgText)
		return
	}

	page.TotalPages = totalPages
	page.DetermineStatus()

	keyboardType := keyboards.ChooseDictionaryKeyboard(page.Status)

	if err := b.uc.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	formatedPage, err := b.uc.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	page.DictionaryMsgID = b.sendMessageWithKeyboard(chatID, formatedPage, keyboardType)
}

func (b *Bot) GetDictionaryCB(userID, chatID int64) {
	userState, err := b.uc.UserStateUC.Get(userID)
	if err != nil || userState == nil {
		logrus.Error(err)
		msgText := b.msg.Errors.UnknownMsg
		b.sendMessage(chatID, msgText)
		return
	}

	page := model.NewDictionaryPage(userID)
	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(page.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		if errMsgText == b.msg.Errors.NoWords {
			page.DictionaryMsgID = b.sendMessageWithKeyboard(chatID, errMsgText, keyboards.AddWordKeyboard)
			return
		}
		b.sendMessage(chatID, errMsgText)
		return
	}

	page.TotalPages = totalPages
	page.DetermineStatus()

	keyboardType := keyboards.ChooseDictionaryKeyboard(page.Status)
	keyboard, ok := keyboardType.(tgbotapi.InlineKeyboardMarkup)
	if !ok {
		b.sendMessage(chatID, b.msg.Errors.Unknown)
		return
	}

	if err := b.uc.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	formatedPage, err := b.uc.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	page.DictionaryMsgID = userState.LastMsgID
	b.updateMessage(chatID, userState.LastMsgID, formatedPage, keyboard)
}

func (b *Bot) GetAnotherDictionaryPage(userID, chatID int64, navigation Navigation) {
	page, err := b.uc.DictionaryPageUC.Get(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	if navigation == Next {
		page.CurrentPage++
	} else {
		page.CurrentPage--
	}

	page.DetermineStatus()
	keyboardType := keyboards.ChooseDictionaryKeyboard(page.Status)

	keyboard, ok := keyboardType.(tgbotapi.InlineKeyboardMarkup)
	if !ok {
		b.sendMessage(chatID, b.msg.Errors.Unknown)
		return
	}

	if err := b.uc.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}

	formatedPage, err := b.uc.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	b.updateMessage(chatID, page.DictionaryMsgID, formatedPage, keyboard)
}
