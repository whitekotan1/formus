package config

import (
	"os"
)

type Config struct {
	Password string
}

func LoadConfig(config *Config) {
	config.Password = os.Getenv("FORMUS_PASSWORD")

}
