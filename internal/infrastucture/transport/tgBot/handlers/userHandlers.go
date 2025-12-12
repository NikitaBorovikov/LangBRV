package handlers

import (
	apperrors "langbrv/internal/app_errors"
	"langbrv/internal/infrastucture/transport/tgBot/dto"

	"github.com/sirupsen/logrus"
)

func (h *Handlers) StartCommand(userID int64, username string) string {
	req := dto.NewRegistrationRequest(userID, username)
	user := req.ToDomainUser()

	if err := h.UseCases.UserUC.CreateOrUpdate(user); err != nil {
		logrus.Error(err)
		errMsgText := apperrors.HandleError(err, &h.Msg.Errors)
		return errMsgText
	}
	return h.Msg.Info.Start
}
