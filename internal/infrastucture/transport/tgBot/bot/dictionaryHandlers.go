package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) GetDictionaryCommand(userID, chatID int64) {
	page := model.NewDictionaryPage(userID)
	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(userID)
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
	page.MessageID = b.sendMessageWithKeyboard(chatID, formatedPage, keyboard)
}

func (b *Bot) GetDictionaryCB(userID, chatID int64) {
	userState, err := b.uc.UserStateUC.Get(userID)
	if err != nil || userState == nil {
		userState = model.NewUserState(userID, false)
		if err := b.uc.UserStateUC.Set(userState); err != nil {
			logrus.Error(err)
			errMsgText := apperrors.HandleError(err, &b.msg.Errors)
			b.sendMessage(chatID, errMsgText)
			return
		}
	}

	page := model.NewDictionaryPage(userID)
	totalPages, err := b.uc.DictionaryPageUC.GetAmountOfPages(page.UserID)
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

	page.MessageID = userState.LastMessageID
	b.updateMessage(chatID, userState.LastMessageID, formatedPage, keyboard)
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

	page.DeterminePosition()
	keyboard := keyboards.ChooseDictionaryKeyboard(page.Position)

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
	b.updateMessage(chatID, page.MessageID, formatedPage, keyboard)
}
