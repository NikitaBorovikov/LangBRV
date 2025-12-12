package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"

	"github.com/sirupsen/logrus"
)

type PageNavigation string

const (
	Next     PageNavigation = "NEXT"
	Previous PageNavigation = "PREVIOUS"
)

func (h *Handlers) GetDictionaryCommand(userID int64) (string, *model.DictionaryPage) {
	page := model.NewDictionaryPage(userID)

	totalPages, err := h.UseCases.DictionaryPageUC.GetAmountOfPages(page.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}
	page.TotalPages = totalPages

	page.DetermineStatus()

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}
	return formatedPage, page
}

func (h *Handlers) GetAnotherDictionaryPage(userID int64, navigation PageNavigation) (string, *model.DictionaryPage) {
	page, err := h.UseCases.DictionaryPageUC.Get(userID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}

	if navigation == Next {
		page.CurrentPage++
	} else {
		page.CurrentPage--
	}

	page.DetermineStatus()

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}
	return formatedPage, page
}
