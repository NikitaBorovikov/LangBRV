package app

import (
	"langbrv/internal/config"
	inmemory "langbrv/internal/infrastucture/repository/inMemory"
	"langbrv/internal/infrastucture/transport/tgBot/bot"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"
	"langbrv/internal/usecases"
	"log"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	inMemoryDB := inmemory.NewUserStateRepo()

	usecases := usecases.NewUseCases(nil, nil, inMemoryDB)
	handlers := handlers.NewHandlers(usecases)

	bot, err := bot.NewBot(&cfg.Telegram, handlers)
	if err != nil {
		log.Fatalf("failed to init telegram bot: %v", err)
	}
	bot.Start()
}
