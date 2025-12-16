package bot

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"

	"github.com/sirupsen/logrus"
)

func (b *Bot) StartCommand(userID, chatID int64, username string) {
	req := dto.NewRegistrationRequest(userID, username)
	user := req.ToDomainUser()

	if err := b.uc.UserUC.CreateOrUpdate(user); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &b.msg.Errors)
		b.sendMessage(chatID, errMsgText)
		return
	}
	msgText := b.msg.Info.Start
	b.sendMessageWithKeyboard(chatID, msgText, keyboards.AddFirstWordKeyboard)
}
