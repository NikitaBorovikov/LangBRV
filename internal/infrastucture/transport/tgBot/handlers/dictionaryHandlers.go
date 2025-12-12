package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/bot/keyboards"

	"github.com/sirupsen/logrus"
)

type PageNavigation string

const (
	Next     PageNavigation = "NEXT"
	Previous PageNavigation = "PREVIOUS"
)

func (h *Handlers) GetDictionaryCommand(userID int64) (string, *model.DictionaryPage, interface{}) {
	page := model.NewDictionaryPage(userID)

	totalPages, err := h.UseCases.DictionaryPageUC.GetAmountOfPages(page.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		if errMsgText == h.Msg.Errors.NoWords {
			return errMsgText, nil, keyboards.AddWordKeyboard
		}
		return errMsgText, nil, nil
	}
	page.TotalPages = totalPages

	page.DetermineStatus()
	keyboardType := keyboards.ChooseDictionaryKeyboard(page.Status)

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil, nil
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil, nil
	}
	return formatedPage, page, keyboardType
}

func (h *Handlers) GetAnotherDictionaryPage(userID int64, navigation PageNavigation) (string, *model.DictionaryPage, interface{}) {
	page, err := h.UseCases.DictionaryPageUC.Get(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil, nil
	}

	if navigation == Next {
		page.CurrentPage++
	} else {
		page.CurrentPage--
	}

	page.DetermineStatus()
	keyboardType := keyboards.ChooseDictionaryKeyboard(page.Status)

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil, nil
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil, nil
	}
	return formatedPage, page, keyboardType
}
