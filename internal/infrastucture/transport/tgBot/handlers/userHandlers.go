package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/infrastucture/transport/tgBot/bot/keyboards"
	"langbrv/internal/infrastucture/transport/tgBot/dto"

	"github.com/sirupsen/logrus"
)

func (h *Handlers) StartCommand(userID int64, username string) (string, interface{}) {
	req := dto.NewRegistrationRequest(userID, username)
	user := req.ToDomainUser()

	if err := h.UseCases.UserUC.CreateOrUpdate(user); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText, nil
	}
	return h.Msg.Info.Start, keyboards.AddFirstWordKeyboard
}
