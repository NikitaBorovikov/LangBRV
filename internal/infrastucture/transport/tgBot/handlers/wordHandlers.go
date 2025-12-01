package handlers

import (
	"fmt"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (h *Handlers) AddWordCommand(update tgbotapi.Update) string {
	var msg string

	state := model.NewUserState(update.Message.Chat.ID, model.AddWord)
	if err := h.UseCases.UserStateUC.Set(state); err != nil {
		msg = "Ошибка"
		return msg
	}

	msg = "Введи слово в формате слово-перевод"
	return msg
}

func (h *Handlers) GetDictionaryCommand(update tgbotapi.Update) string {
	words, err := h.UseCases.WordUC.GetAll(update.Message.From.ID)
	if err != nil {
		logrus.Error(err)
		return "Ошибка"
	}

	dictionary, err := h.UseCases.WordUC.FormatDictionary(words)
	if err != nil {
		logrus.Error(err)
		return "Ошибка"
	}
	return dictionary
}

func (h *Handlers) SaveWord(update tgbotapi.Update) string {
	req := dto.NewAddWordRequest(update.Message.From.ID, update.Message.Text)
	word, err := req.ToDomainWord()
	if err != nil {
		return "Некорректный формат ввода"
	}

	wordID, err := h.UseCases.WordUC.Add(word)
	if err != nil {
		log.Printf("error: %v", err)
		return "Ошибка сохранения"
	}
	fmt.Printf("word is saved: id = %s", wordID)
	return "Слово сохранено"
}
