package config

import (
	"os"
)

type Config struct {
	Password string
}

func LoadConfig() *Config {

	return &Config{
		Password: os.Getenv("FORMUS_PASSWORD"),
	}

}

func ValidatePassword(receivedPassword string, cfg *Config) bool {
	if receivedPassword == cfg.Password {
		return true
	}
	return false
}
