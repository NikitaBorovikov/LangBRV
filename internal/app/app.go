package app

import (
	"fmt"
	"langbrv/internal/config"
	repo "langbrv/internal/infrastucture/repository"
	"langbrv/internal/infrastucture/transport/tgBot/bot"
	"langbrv/internal/infrastucture/transport/tgBot/handlers"
	"langbrv/internal/usecases"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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

	repo := repo.NewRepository(db)
	usecases := usecases.NewUseCases(repo)
	handlers := handlers.NewHandlers(usecases, &cfg.Msg)

	bot, err := bot.NewBot(&cfg.Telegram, handlers)
	if err != nil {
		logrus.Fatalf("failed to init Telegram bot: %v", err)
	}
	bot.Start()
}

func initPostgresDB(cfg *config.DB) (*sqlx.DB, error) {
	dsn := makeDSN(cfg)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}
	return db, err
}

func makeDSN(cfg *config.DB) string {
	return fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", cfg.User, cfg.Name, cfg.Password, cfg.Host, cfg.Port,
	)
}
