package bot

import (
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"
)

func (b *Bot) StartCommand(userID, chatID int64, username string) {
	user := model.NewUser(userID, username)

	if err := b.uc.UserUC.CreateOrUpdate(user); err != nil {
		b.handleError(chatID, err)
		return
	}
	msgText := b.msg.Info.Start
	b.sendMessageWithKeyboard(chatID, msgText, keyboards.StartKeyboard)
}
