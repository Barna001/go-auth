package config

import (
	"github.com/Barna001/go-auth/errors"
	"github.com/joho/godotenv"
)

type AuthConfig struct {
	TextDBLocation string `default:"users.txt"`
	WebserverPort  int    `default:"8080"`
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		errors.CriticalHandling(err)
	}
}