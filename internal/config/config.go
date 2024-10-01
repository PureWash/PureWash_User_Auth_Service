package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	USER_SERVER string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SECRET_KEY  string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}

	var cfg Config

	cfg.USER_SERVER = cast.ToString(coalesce("USER_SERVICE", "localhost:8081"))

	cfg.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(coalesce("DB_PORT", "5432"))
	cfg.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	cfg.DB_NAME = cast.ToString(coalesce("DB_NAME", "postgres"))
	cfg.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "password"))

	cfg.SECRET_KEY = cast.ToString(coalesce("SECRET_KEY", "my_secret_key"))

	return cfg
}

func coalesce(key string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultValue
}
