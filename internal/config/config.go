package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configYAMLPath = "config/config.yaml"
	configENVPath  = ".env"
)

type Config struct {
	Telegram Telegram `yaml:"telegram"`
	DB       DB
	Msg      Messages `yaml:"messages"`
}

type Telegram struct {
	Token         string `env:"TELEGRAM_TOKEN"`
	UpdateTimeout int    `yaml:"update_timeout"`
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
	DelWord string `yaml:"del_word"`
}

type Success struct {
	WordAdded   string `yaml:"word_added"`
	WordDeleted string `yaml:"word_deleted"`
}

type Errors struct {
	Unknown         string `yaml:"unknown"`
	UnknownCommand  string `yaml:"unknown_command"`
	UnknownMsg      string `yaml:"unknown_msg"`
	NoWords         string `yaml:"no_words"`
	NoWordsToRemind string `yaml:"no_words_to_remind"`
	WordNotExists   string `yaml:"word_not_exists"`
	WordTooLong     string `yaml:"word_too_long"`
	IncorrectFormat string `yaml:"incorrect_fomat"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configYAMLPath, &cfg); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadConfig(configENVPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
