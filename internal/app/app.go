package app

import (
	"langbrv/internal/config"
	repo "langbrv/internal/infrastucture/repository"
	"langbrv/internal/infrastucture/transport/tgBot/bot"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"
	"langbrv/internal/usecases"
	"log"

	"gorm.io/gorm"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	db := &gorm.DB{} //REMOVE
	repo := repo.NewRepository(db)
	usecases := usecases.NewUseCases(repo)
	handlers := handlers.NewHandlers(usecases)

	bot, err := bot.NewBot(&cfg.Telegram, handlers)
	if err != nil {
		log.Fatalf("failed to init telegram bot: %v", err)
	}
	bot.Start()
}
