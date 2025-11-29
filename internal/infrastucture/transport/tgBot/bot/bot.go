package bot

import (
	"langbrv/internal/config"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		// REMOVE: for testing
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := b.bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
