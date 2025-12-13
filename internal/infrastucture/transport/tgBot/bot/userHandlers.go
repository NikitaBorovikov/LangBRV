package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (b *Bot) StartCommand(msg tgbotapi.Message) {
	req := dto.NewRegistrationRequest(msg.From.ID, msg.From.UserName)
	user := req.ToDomainUser()

	if err := b.uc.UserUC.CreateOrUpdate(user); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.cfg.Msg.Errors)
		b.sendMessage(msg.Chat.ID, errMsgText)
		return
	}
	msgText := b.cfg.Msg.Info.Start
	b.sendMessageWithKeyboard(msg.Chat.ID, msgText, keyboards.AddFirstWordKeyboard)
}
