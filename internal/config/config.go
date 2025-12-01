package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Telegram Telegram
	DB       DB
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN"`
}

type DB struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Port     int    `env:"DB_PORT"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg.Telegram); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg.DB); err != nil {
		return nil, err
	}
	return &cfg, nil
}
