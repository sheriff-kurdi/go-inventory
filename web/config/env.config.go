package config

import (
	"log"

	"github.com/joho/godotenv"
)

func ENVConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
