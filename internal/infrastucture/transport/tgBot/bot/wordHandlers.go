package bot

import (
	"fmt"
	"langbrv/internal/core/model"
	"langbrv/internal/infrastucture/transport/tgBot/dto"
	"langbrv/internal/infrastucture/transport/tgBot/keyboards"
)

func (b *Bot) AddWord(us *model.UserState, chatID int64) {
	us.IsDeleteMode = false
	msgText := b.msg.Info.AddWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) DeleteWordCommand(us *model.UserState, chatID int64) {
	us.IsDeleteMode = true
	if err := b.uc.UserStateUC.Save(us); err != nil {
		b.handleError(chatID, err)
		return
	}
	msgText := b.msg.Info.DelWord
	b.sendMessage(chatID, msgText)
}

func (b *Bot) SaveWord(us *model.UserState, chatID int64, text string) {
	req := dto.NewAddWordRequest(us.UserID, text)
	word, err := req.ToDomainWord()
	if err != nil {
		b.handleError(chatID, err)
		return
	}

	if err := b.uc.WordUC.Add(word); err != nil {
		b.handleError(chatID, err)
		return
	}

	msgText := b.msg.Success.WordAdded
	msgID := b.sendMessageWithKeyboard(chatID, msgText, keyboards.MainKeyboard)
	us.LastMessageID = msgID
}

func (b *Bot) DeleteWord(us *model.UserState, chatID int64, text string) {
	defer func() {
		us.IsDeleteMode = false
	}()

	amountOfDeletedWords, err := b.uc.WordUC.Delete(us.UserID, text)
	if err != nil || amountOfDeletedWords == 0 {
		b.handleError(chatID, err)
		return
	}

	msgText := b.msg.Success.WordDeleted + fmt.Sprintf("%d", amountOfDeletedWords)
	msgID := b.sendMessageWithKeyboard(chatID, msgText, keyboards.MainKeyboard)
	us.LastMessageID = msgID
}
