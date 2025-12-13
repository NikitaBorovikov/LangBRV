package bot

import (
	"langbrv/internal/config"
	"langbrv/internal/core/model"
	"langbrv/internal/usecases"
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

	NextPageCallback      = "nextPage"
	PreviousPageCallback  = "previousPage"
	AddWordCallback       = "addWord"
	GetDictionaryCallback = "getDictionary"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	cfg *config.Config
	uc  *usecases.UseCases
}

func NewBot(cfg *config.Config, uc *usecases.UseCases) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true

	b := &Bot{
		bot: bot,
		cfg: cfg,
		uc:  uc,
	}
	return b, nil
}

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = b.cfg.Telegram.UpdateTimeout
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
		b.StartCommand(*update.Message)

	case AddWordCommand:
		b.AddWordCommand(update.Message.From.ID, update.Message.Chat.ID)

	case GetDictionaryCommand:
		b.GetDictionaryCommand(update.Message.From.ID, update.Message.Chat.ID)

	case RemindCommand:
		b.GetRemindListCommand(update)

	case DeleteWordCommand:
		b.DeleteWordCommand(update)

	default:
		msgText := b.cfg.Msg.Errors.UnknownCommand
		b.sendMessage(update.Message.Chat.ID, msgText)
	}
}

func (b *Bot) handleMessages(update tgbotapi.Update) {
	userState, err := b.uc.UserStateUC.Get(update.Message.From.ID)
	if err != nil || userState == nil {
		logrus.Error(err)
		msgText := b.cfg.Msg.Errors.UnknownMsg
		b.sendMessage(update.Message.Chat.ID, msgText)
		return
	}

	switch userState.State {
	case model.AddWord:
		b.SaveWord(update)

	case model.DelWord:
		b.DeleteWord(update)

	default:
		msgText := b.cfg.Msg.Errors.UnknownMsg
		b.sendMessage(update.Message.Chat.ID, msgText)
	}
}

func (b *Bot) handleCallbacks(update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case NextPageCallback:
		b.GetAnotherDictionaryPage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.Chat.ID, Next)

	case PreviousPageCallback:
		b.GetAnotherDictionaryPage(update.CallbackQuery.From.ID, update.CallbackQuery.Message.Chat.ID, Previous)

	case AddWordCallback:
		b.AddWordCommand(update.CallbackQuery.From.ID, update.CallbackQuery.Message.Chat.ID)

	case GetDictionaryCallback:
		b.GetDictionaryCommand(update.CallbackQuery.From.ID, update.CallbackQuery.Message.Chat.ID)

	default:
		msgText := b.cfg.Msg.Errors.Unknown
		b.sendMessage(update.CallbackQuery.Message.Chat.ID, msgText)
	}
}

func (b *Bot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	if _, err := b.bot.Send(msg); err != nil {
		logrus.Errorf("failed to send message to chat id: %d, err: %v", chatID, err)
	}
}

func (b *Bot) sendMessageWithKeyboard(chatID int64, text string, keyboard interface{}) int {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = keyboard
	msgInfo, err := b.bot.Send(msg)
	if err != nil {
		logrus.Errorf("failed to send message to chat id: %d, err: %v", chatID, err)
		return 0
	}
	return msgInfo.MessageID
}

func (b *Bot) updateMessage(chatID int64, msgID int, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewEditMessageText(chatID, msgID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = &keyboard
	if _, err := b.bot.Send(msg); err != nil {
		logrus.Errorf("failed to update message to chat id: %d, err: %v", chatID, err)
	}
}
