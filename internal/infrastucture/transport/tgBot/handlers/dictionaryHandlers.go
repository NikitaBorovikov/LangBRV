package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type DictionaryPageStatus string

const (
	SinglePage DictionaryPageStatus = "SINGLE"
	FirstPage  DictionaryPageStatus = "FIRST"
	MiddlePage DictionaryPageStatus = "MIDDLE"
	LastPage   DictionaryPageStatus = "LAST"
)

type PageNavigation string

const (
	Next     PageNavigation = "NEXT"
	Previous PageNavigation = "PREVIOUS"
)

func (h *Handlers) GetDictionaryCommand(update tgbotapi.Update) (string, DictionaryPageStatus) {
	page := model.NewDictionaryPage(update.Message.From.ID)

	totalPages, err := h.UseCases.DictionaryPageUC.GetAmountOfPages(page.UserID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, ""
	}
	page.TotalPages = totalPages

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, ""
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, ""
	}
	pageStatus := DeterminePageStatus(page.CurrentPage, page.TotalPages)
	return formatedPage, pageStatus
}

func (h *Handlers) GetAnotherDictionaryPage(update tgbotapi.Update, navigation PageNavigation) (string, DictionaryPageStatus, int) {
	page, err := h.UseCases.DictionaryPageUC.Get(update.CallbackQuery.From.ID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, "", 0
	}

	if navigation == Next {
		page.CurrentPage++
	} else {
		page.CurrentPage--
	}

	if err := h.UseCases.DictionaryPageUC.Save(page); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, "", 0
	}

	formatedPage, err := h.UseCases.DictionaryPageUC.FormatPage(page)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, "", 0
	}
	pageStatus := DeterminePageStatus(page.CurrentPage, page.TotalPages)
	return formatedPage, pageStatus, page.DictionaryMsgID
}

func DeterminePageStatus(currentPage, totalPages int64) DictionaryPageStatus {
	switch {
	case currentPage == 1 && totalPages == 1:
		return SinglePage
	case currentPage == 1 && totalPages > 1:
		return FirstPage
	case currentPage != 1 && currentPage == totalPages:
		return LastPage
	default:
		return MiddlePage
	}
}
