package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver		string
	DBName			string
	DBHost			string
	DBPort			string
	DBUser			string
	DBPassword		string

	TelegramToken	string
	AppPort			string
	LogLevel		string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env")
	}
	cfg := &Config{
		DBDriver: 		getenv("DB_DRIVER", "sqlite3"),
		DBName: 		getenv("DB_NAME", "tasktracker.db"),
		DBHost: 		getenv("DB_HOST", "localhost"),
		DBPort:			getenv("DB_PORT", "5432"),
		DBUser:        	getenv("DB_USER", ""),
		DBPassword:    	getenv("DB_PASSWORD", ""),
		AppPort:       	getenv("APP_PORT", "8080"),
		TelegramToken: 	getenv("TELEGRAM_TOKEN", ""),
		LogLevel:      	getenv("LOG_LEVEL", "debug"),
	}

	return cfg
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}