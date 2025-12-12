package bot

import (
	"langbrv/internal/config"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

const (
	StartCommand         = "start"
	AddWordCommand       = "add"
	GetDictionaryCommand = "dictionary"
	RemindCommand        = "remind"
	DeleteWordCommand    = "del_word"

	NextPageCallback     = "nextPage"
	PreviousPageCallback = "previousPage"
	AddWordCallback      = "addWord"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	cfg      *config.Telegram
	handlers *handlers.Handlers
}

func NewBot(cfg *config.Telegram, handlers *handlers.Handlers) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true

	b := &Bot{
		bot:      bot,
		cfg:      cfg,
		handlers: handlers,
	}
	return b, nil
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = b.cfg.UpdateTimeout
	updates := b.bot.GetUpdatesChan(updateConfig)
	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			start := time.Now()
			if update.Message.IsCommand() {
				b.handleCommands(update)
			} else {
				b.handleMessages(update)
			}
			duration := time.Since(start)
			logrus.Infof("Request duration: %v", duration)

		} else if update.CallbackQuery != nil {
			start := time.Now()
			b.handleCallbacks(update)
			duration := time.Since(start)
			logrus.Infof("Request duration: %v", duration)
		}
	}
}

func (b *Bot) handleCommands(update tgbotapi.Update) {
	switch update.Message.Command() {
	case StartCommand:
		msgText := b.handlers.StartCommand(update.Message.From.ID, update.Message.From.UserName)
		b.sendMessageWithKeyboard(update.Message.Chat.ID, msgText, AddFirstWordKeyboard)

	case AddWordCommand:
		msgText := b.handlers.AddWordCommand(update.Message.From.ID)
		b.sendMessage(update.Message.Chat.ID, msgText)

	case GetDictionaryCommand:
		msgText, pageInfo := b.handlers.GetDictionaryCommand(update.Message.From.ID)
		if pageInfo != nil {
			keyboard := ChooseDictionaryKeyboard(pageInfo.Status)
			msgID := b.sendMessageWithKeyboard(update.Message.Chat.ID, msgText, keyboard)
			pageInfo.DictionaryMsgID = msgID
			return
		}
		if msgText == b.handlers.Msg.Errors.NoWords {
			b.sendMessageWithKeyboard(update.Message.Chat.ID, msgText, AddWordKeyboard)
			return
		}
		b.sendMessage(update.Message.Chat.ID, msgText)

	case RemindCommand:
		msgText := b.handlers.GetRemindListCommand(update.Message.From.ID)
		b.sendMessage(update.Message.Chat.ID, msgText)

	case DeleteWordCommand:
		msgText := b.handlers.DeleteWordCommand(update.Message.From.ID)
		b.sendMessage(update.Message.Chat.ID, msgText)

	default:
		msgText := b.handlers.Msg.Errors.UnknownCommand
		b.sendMessage(update.Message.Chat.ID, msgText)
	}
}

func (b *Bot) handleMessages(update tgbotapi.Update) {
	userState, err := b.handlers.UseCases.UserStateUC.Get(update.Message.From.ID)
	if err != nil || userState == nil {
		logrus.Error(err)
		msgText := b.handlers.Msg.Errors.UnknownMsg
		b.sendMessage(update.Message.Chat.ID, msgText)
		return
	}

	switch userState.State {
	case model.AddWord:
		msgText := b.handlers.SaveWord(update.Message.From.ID, update.Message.Text)
		b.sendMessage(update.Message.Chat.ID, msgText)

	case model.DelWord:
		msgText := b.handlers.DeleteWord(update.Message.From.ID, update.Message.Text)
		b.sendMessage(update.Message.Chat.ID, msgText)

	default:
		msgText := b.handlers.Msg.Errors.UnknownMsg
		b.sendMessage(update.Message.Chat.ID, msgText)
	}
}

func (b *Bot) handleCallbacks(update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case NextPageCallback:
		msgText, pageInfo := b.handlers.GetAnotherDictionaryPage(update.CallbackQuery.From.ID, handlers.Next)
		keyboard := ChooseDictionaryKeyboard(pageInfo.Status)
		b.updateDictionaryMsg(update.CallbackQuery.Message.Chat.ID, pageInfo.DictionaryMsgID, msgText, keyboard)

	case PreviousPageCallback:
		msgText, pageInfo := b.handlers.GetAnotherDictionaryPage(update.CallbackQuery.From.ID, handlers.Previous)
		keyboard := ChooseDictionaryKeyboard(pageInfo.Status)
		b.updateDictionaryMsg(update.CallbackQuery.Message.Chat.ID, pageInfo.DictionaryMsgID, msgText, keyboard)

	case AddWordCallback:
		msgText := b.handlers.AddWordCommand(update.CallbackQuery.From.ID)
		b.sendMessage(update.CallbackQuery.Message.Chat.ID, msgText)

	default:
		msgText := b.handlers.Msg.Errors.Unknown
		b.sendMessage(update.CallbackQuery.Message.Chat.ID, msgText)
	}
}

func (b *Bot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := b.bot.Send(msg); err != nil {
		logrus.Errorf("failed to send message to chat id: %d, err: %v", chatID, err)
	}
}

func (b *Bot) sendMessageWithKeyboard(chatID int64, text string, keyboard tgbotapi.InlineKeyboardMarkup) int {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	msgInfo, err := b.bot.Send(msg)
	if err != nil {
		logrus.Errorf("failed to send message to chat id: %d, err: %v", chatID, err)
		return 0
	}
	return msgInfo.MessageID
}

func (b *Bot) updateDictionaryMsg(chatID int64, msgID int, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewEditMessageText(chatID, msgID, text)
	msg.ReplyMarkup = &keyboard
	if _, err := b.bot.Send(msg); err != nil {
		logrus.Errorf("failed to update message to chat id: %d, err: %v", chatID, err)
	}
}
