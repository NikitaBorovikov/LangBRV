package bot

import (
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"
)

func (b *Bot) StartCommand(userID, chatID int64, username string) {
	req := dto.NewRegistrationRequest(userID, username)
	user := req.ToDomainUser()

	if err := b.uc.UserUC.CreateOrUpdate(user); err != nil {
		b.handleError(chatID, err)
		return
	}
	msgText := b.msg.Info.Start
	b.sendMessageWithKeyboard(chatID, msgText, keyboards.StartKeyboard)
}
