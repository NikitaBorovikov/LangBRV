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

func (h *Handlers) GetDictionaryCommand(update tgbotapi.Update) (string, DictionaryPageStatus) {
	page := model.NewDictionaryPage(update.Message.From.ID)

	totalPages, err := h.UseCases.WordUC.GetAmountOfPages(page.UserID)
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
