package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	PostgresDB 	string
	PostgresUser string
	PostgresPassword string
	PostgresPort string
}

func Load()	Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("warning: .env file not found")
	}

	cfg := Config {
		PostgresDB: os.Getenv("POSTGRES_DB"),
		PostgresUser:	os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresPort: os.Getenv("POSTGRES_PORT"),
	}

	return cfg     // Get the database variables
}


