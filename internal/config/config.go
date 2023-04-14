package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	APIHost string
	APIPort string

	DEBUG bool
}

func LoadConfig(dotenv_path string) Config {
	err := godotenv.Load(dotenv_path)

	if err != nil {
		log.Fatal("Error while loading .env file: ", err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Println("Error while parsing DEBUG to bool")
		debug = false
	}

	return Config{
		DBUsername: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASS"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),

		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),

		DEBUG: debug,
	}
}
