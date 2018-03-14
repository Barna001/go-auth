package config

import (
	"github.com/Barna001/go-auth/errors"
	"github.com/joho/godotenv"
)

type AuthConfig struct {
	TextDBLocation string `default:"users.txt"`
	WebserverPort  int    `default:"8080"`
	JWTSignKey     string `required:"true"`
}

func LoadEnv() {
	err := godotenv.Load()
	errors.CriticalHandling(err)
}
