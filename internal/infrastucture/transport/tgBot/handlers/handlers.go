package handlers

import (
	"langbrv/internal/config"
	"langbrv/internal/usecases"
)

type Handlers struct {
	UseCases *usecases.UseCases
	Msg      *config.Messages
}

func NewHandlers(uc *usecases.UseCases, msg *config.Messages) *Handlers {
	return &Handlers{
		UseCases: uc,
		Msg:      msg,
	}
}
