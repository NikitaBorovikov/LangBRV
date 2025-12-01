package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	configPath = "config/config.yaml"
)

type Config struct {
	Telegram Telegram
	DB       DB
	Msg      Messages `yaml:"messages"`
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

type Messages struct {
	Info    Info    `yaml:"info"`
	Success Success `yaml:"success"`
	Errors  Errors  `yaml:"errors"`
}

type Info struct {
	Start   string `yaml:"start"`
	AddWord string `yaml:"add_word"`
}

type Success struct {
	WordAdded   string `yaml:"word_added"`
	WordDeleted string `yaml:"word_deleted"`
}

type Errors struct {
	Unknown        string `yaml:"unknown"`
	UnknownCommand string `yaml:"unknown_command"`
	NoWords        string `yaml:"no_words"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
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
