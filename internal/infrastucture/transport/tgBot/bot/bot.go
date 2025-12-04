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
	AddWordCommand       = "add_word"
	GetDictionaryCommand = "dictionary"
	RemindCommand        = "remind"
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
		if update.Message == nil {
			continue
		}

		start := time.Now()

		if update.Message.IsCommand() {
			b.handleCommands(update)
		} else {
			b.handleMessages(update)
		}

		duration := time.Since(start)
		logrus.Infof("Request duration: %v", duration)
	}
}

func (b *Bot) handleCommands(update tgbotapi.Update) {
	switch update.Message.Command() {

	case StartCommand:
		msgText := b.handlers.StartCommand(update)
		b.sendMessage(update, msgText)

	case AddWordCommand:
		msgText := b.handlers.AddWordCommand(update)
		b.sendMessage(update, msgText)

	case GetDictionaryCommand:
		msgText := b.handlers.GetDictionaryCommand(update)
		b.sendMessage(update, msgText)

	case RemindCommand:
		msgTest := b.handlers.GetRemindListCommand(update)
		b.sendMessage(update, msgTest)

	default:
		msgText := b.handlers.Msg.Errors.UnknownCommand
		b.sendMessage(update, msgText)
	}
}

func (b *Bot) handleMessages(update tgbotapi.Update) {
	userState, err := b.handlers.UseCases.UserStateUC.Get(update.Message.From.ID)
	if err != nil || userState == nil {
		logrus.Error(err)
		msgText := b.handlers.Msg.Errors.UnknownMsg
		b.sendMessage(update, msgText)
		return
	}

	switch userState.State {

	case model.AddWord:
		msgText := b.handlers.SaveWord(update)
		b.sendMessage(update, msgText)

	default:
		msgText := b.handlers.Msg.Errors.UnknownMsg
		b.sendMessage(update, msgText)
	}
}

func (b *Bot) sendMessage(update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := b.bot.Send(msg); err != nil {
		logrus.Errorf("failed to send message to chat id: %d, err: %v", update.Message.Chat.ID, err)
	}
}
