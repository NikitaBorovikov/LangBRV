package bot

import (
	"langbrv/internal/config"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	StartCommand   = "start"
	AddWordCommand = "add_word"
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
	updateConfig.Timeout = 30
	updates := b.bot.GetUpdatesChan(updateConfig)

	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommands(update)
		}
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
	}
}

func (b *Bot) sendMessage(update tgbotapi.Update, text string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := b.bot.Send(msg); err != nil {
		log.Printf("failed to send message to chat id: %d, err: %v", update.Message.Chat.ID, err)
	}
}
