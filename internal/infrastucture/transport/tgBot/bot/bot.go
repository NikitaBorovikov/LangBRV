package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/config"
	"langbrv/internal/core/model"
	"langbrv/internal/usecases"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Navigation string

const (
	Next     Navigation = "NEXT"
	Previous Navigation = "PREVIOUS"
)

const (
	StartCommand         = "start"
	AddWordCommand       = "add"
	GetDictionaryCommand = "dictionary"
	RemindCommand        = "remind"
	DeleteWordCommand    = "delete"

	NextPageCallback      = "nextPage"
	PreviousPageCallback  = "previousPage"
	AddWordCallback       = "addWord"
	GetDictionaryCallback = "getDictionary"
	RemindSessionCallback = "newRemindSession"

	RememberWellCallback  = "rememberWell"
	RememberBadlyCallback = "rememberBadly"

	ShowWordCallback = "showWord"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	msg *config.Messages
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
		msg: &cfg.Msg,
		uc:  uc,
	}
	return b, nil
}

func (b *Bot) Start(cfg *config.Telegram) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = cfg.UpdateTimeout
	updates := b.bot.GetUpdatesChan(updateConfig)
	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		go func(u tgbotapi.Update) {
			defer func() {
				if r := recover(); r != nil {
					logrus.Errorf("recovered from panic: %v", r)
				}
			}()

			b.processUpdate(u)
		}(update)
	}
}

func (b *Bot) processUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		start := time.Now()
		if update.Message.IsCommand() {
			b.handleCommands(update.Message)
		} else {
			b.handleMessages(update.Message)
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

func (b *Bot) handleCommands(update *tgbotapi.Message) {
	userID := update.From.ID
	chatID := update.Chat.ID

	userState, err := b.uc.UserStateUC.Get(userID)
	if err != nil || userState == nil {
		userState = model.NewUserState(userID)
		if err := b.uc.UserStateUC.Save(userState); err != nil {
			logrus.Error(err)
			errMsgText := apperrors.HandleError(err, &b.msg.Errors)
			b.sendMessage(chatID, errMsgText)
			return
		}
	}

	switch update.Command() {
	case StartCommand:
		username := update.From.UserName
		b.StartCommand(userID, chatID, username)

	case AddWordCommand:
		b.AddWord(userState, chatID)

	case GetDictionaryCommand:
		b.GetDictionaryCommand(userState, chatID)

	case RemindCommand:
		b.StartRemindSession(userState, chatID)

	case DeleteWordCommand:
		b.DeleteWordCommand(userState, chatID)

	default:
		msgText := b.msg.Errors.UnknownCommand
		b.sendMessage(chatID, msgText)
	}
}

func (b *Bot) handleMessages(update *tgbotapi.Message) {
	userID := update.From.ID
	chatID := update.Chat.ID
	text := update.Text

	userState, err := b.uc.UserStateUC.Get(userID)
	if err != nil || userState == nil {
		userState = model.NewUserState(userID)
		if err := b.uc.UserStateUC.Save(userState); err != nil {
			logrus.Error(err)
			errMsgText := apperrors.HandleError(err, &b.msg.Errors)
			b.sendMessage(chatID, errMsgText)
			return
		}
	}

	if !userState.IsDeleteMode {
		b.SaveWord(userState, chatID, text)
	} else {
		b.DeleteWord(userState, chatID, text)
	}
}

func (b *Bot) handleCallbacks(update tgbotapi.Update) {
	userID := update.CallbackQuery.From.ID
	chatID := update.CallbackQuery.Message.Chat.ID

	userState, err := b.uc.UserStateUC.Get(userID)
	if err != nil || userState == nil {
		userState = model.NewUserState(userID)
		if err := b.uc.UserStateUC.Save(userState); err != nil {
			logrus.Error(err)
			errMsgText := apperrors.HandleError(err, &b.msg.Errors)
			b.sendMessage(chatID, errMsgText)
			return
		}
	}

	switch update.CallbackQuery.Data {
	case NextPageCallback:
		b.GetAnotherDictionaryPage(userState, chatID, Next)

	case PreviousPageCallback:
		b.GetAnotherDictionaryPage(userState, chatID, Previous)

	case ShowWordCallback:
		b.ShowRemindCard(userState, chatID)

	case RememberWellCallback:
		b.GetNextRemindCard(userState, chatID, true)

	case RememberBadlyCallback:
		b.GetNextRemindCard(userState, chatID, false)

	case AddWordCallback:
		b.AddWord(userState, chatID)

	case GetDictionaryCallback:
		b.GetDictionaryCB(userState, chatID)

	case RemindSessionCallback:
		b.RepeatRemindSession(userState, chatID)

	default:
		msgText := b.msg.Errors.Unknown
		b.sendMessage(chatID, msgText)
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

func (b *Bot) updateMessage(chatID int64, msgID int, text string, keyboardType interface{}) {
	msg := tgbotapi.NewEditMessageText(chatID, msgID, text)
	msg.ParseMode = tgbotapi.ModeHTML
	keyboard, ok := keyboardType.(tgbotapi.InlineKeyboardMarkup)
	if !ok {
		if _, err := b.bot.Send(msg); err != nil {
			logrus.Errorf("failed to update message to chat id: %d, err: %v", chatID, err)
		}
		return
	}
	msg.ReplyMarkup = &keyboard
	if _, err := b.bot.Send(msg); err != nil {
		logrus.Errorf("failed to update message to chat id: %d, err: %v", chatID, err)
	}
}
