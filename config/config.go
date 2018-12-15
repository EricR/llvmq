package config

import (
	"log"
	"os"
)

type Config struct {
	Logger *log.Logger
}

func New() *Config {
	return &Config{
		Logger: log.New(os.Stderr, "", log.Ldate|log.Ltime),
	}
}
