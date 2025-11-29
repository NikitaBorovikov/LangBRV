package app

import (
	"langbrv/internal/config"
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

	usecases := usecases.NewUseCases(nil, nil)
	handlers := handlers.NewHandlers(usecases)

	bot, err := bot.NewBot(&cfg.Telegram, handlers)
	if err != nil {
		log.Fatalf("failed to init telegram bot: %v", err)
	}
	bot.Start()
}
