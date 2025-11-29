package app

import (
	"fmt"
	"langbrv/internal/config"
	repo "langbrv/internal/infrastucture/repository"
	"langbrv/internal/infrastucture/transport/tgBot/bot"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"
	"langbrv/internal/usecases"

	"github.com/sirupsen/logrus"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("failed to init config: %v", err)
	}

	db, err := initPostgresDB(&cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to init PostgreSQL: %v", err)
	}

	// TODO: add auto migrate

	repo := repo.NewRepository(db)
	usecases := usecases.NewUseCases(repo)
	handlers := handlers.NewHandlers(usecases)

	bot, err := bot.NewBot(&cfg.Telegram, handlers)
	if err != nil {
		logrus.Fatalf("failed to init Telegram bot: %v", err)
	}
	bot.Start()
}

func initPostgresDB(cfg *config.DB) (*gorm.DB, error) {
	dsn := makeDSN(cfg)
	db, err := gorm.Open(p.Open(dsn))
	return db, err
}

func makeDSN(cfg *config.DB) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
}
