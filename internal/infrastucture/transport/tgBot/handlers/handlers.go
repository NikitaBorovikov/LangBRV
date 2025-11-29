package handlers

import "langbrv/internal/usecases"

type Handlers struct {
	UseCases *usecases.UseCases
}

func NewHandlers(uc *usecases.UseCases) *Handlers {
	return &Handlers{UseCases: uc}
}
