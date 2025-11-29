package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Telegram Telegram
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg.Telegram); err != nil {
		return nil, err
	}
	return &cfg, nil
}
