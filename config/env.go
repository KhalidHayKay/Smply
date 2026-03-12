package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type EnvType struct {
	AppEnv  string
	AppPort string
	AppUrl  string
	DbUrl   string
}

var Env *EnvType

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	Env = &EnvType{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: os.Getenv("PORT"),
		AppUrl:  os.Getenv("APP_URL"),
		DbUrl:   os.Getenv("DB_URL"),
	}

	if Env.AppUrl == "" || Env.DbUrl == "" {
		return errors.New("APP_URL or DB_URL not set")
	}

	return nil
}
