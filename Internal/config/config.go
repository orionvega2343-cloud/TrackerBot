package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DB struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
	Password string `env:"DB_PASS"`
	SslMode  string `yaml:"ssl_mode"`
}

type Config struct {
	Token    string `env:"BOT_TOKEN"`
	ProxyURL string `env:"PROXY_URL"`
	DB       DB     `yaml:"db"`
}

func MustLoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	var c Config

	err = cleanenv.ReadConfig("config.yml", &c)
	if err != nil {
		log.Fatal(err)
	}
	return &c
}
