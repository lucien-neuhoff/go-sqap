package config

import (
	"log"
	"os"
	"path/filepath"
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

	PREFIX_PATH string
}

func LoadConfig(dotenv_path string) Config {
	err := godotenv.Load(dotenv_path)

	if err != nil {
		log.Fatal("Error while loading .env file: ", err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		log.Println("Error while parsing DEBUG to bool: ", err)
		debug = false
	}

	workingDir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Println("Error while getting working directory: ", err)
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

		PREFIX_PATH: workingDir,
	}
}
