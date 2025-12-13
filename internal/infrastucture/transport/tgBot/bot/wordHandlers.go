package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *Bot) AddWordCommand(userID, chatID int64) {
	state := model.NewUserState(userID, model.AddWord, 0)

	if err := b.uc.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.cfg.Msg.Info.AddWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) GetRemindListCommand(update tgbotapi.Update) {
	remindList, err := b.uc.WordUC.GetRemindList(update.Message.From.ID)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(update.Message.Chat.ID, errMsgText)
		return
	}

	remindMsg, err := b.uc.WordUC.FormatRemindList(remindList)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(update.Message.Chat.ID, errMsgText)
		return
	}
	b.sendMessage(update.Message.Chat.ID, remindMsg)
}

func (b *Bot) DeleteWordCommand(update tgbotapi.Update) {
	state := model.NewUserState(update.Message.From.ID, model.DelWord, 0)

	if err := b.uc.UserStateUC.Set(state); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(update.Message.Chat.ID, errMsgText)
		return
	}
	msgText := b.cfg.Msg.Info.DelWord
	b.sendMessage(update.Message.Chat.ID, msgText)
}

func (b *Bot) SaveWord(update tgbotapi.Update) {
	userState, err := b.uc.UserStateUC.Get(update.Message.From.ID)
	if err != nil || userState == nil {
		logrus.Error(err)
		msgText := b.cfg.Msg.Errors.UnknownMsg
		b.sendMessage(update.Message.Chat.ID, msgText)
		return
	}

	req := dto.NewAddWordRequest(update.Message.From.ID, update.Message.Text)
	word, err := req.ToDomainWord()
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(update.Message.Chat.ID, errMsgText)
		return
	}

	wordID, err := b.uc.WordUC.Add(word)
	if err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(update.Message.Chat.ID, errMsgText)
		return
	}

	msgText := b.cfg.Msg.Success.WordAdded
	msgID := b.sendMessageWithKeyboard(update.Message.Chat.ID, msgText, keyboards.MainKeyboard)
	userState.LastMsgID = msgID
	logrus.Infof("word is saved: id = %s", wordID)
}

func (b *Bot) DeleteWord(update tgbotapi.Update) {
	if err := b.uc.WordUC.Delete(update.Message.From.ID, update.Message.Text); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(update.Message.Chat.ID, errMsgText)
		return
	}
	msgText := b.cfg.Msg.Success.WordDeleted
	b.sendMessage(update.Message.Chat.ID, msgText)
}
