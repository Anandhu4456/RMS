package db

import (
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_PORT",
}

func LoadConfig() (Config, error) {

	var config Config

	config.DBHost = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_NAME")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBPort = os.Getenv("DB_PORT")

	for _, env := range envs {
		if value := os.Getenv(env); value == "" {
			log.Fatalf("[missing env variable]: %s is not set", env)
		}
	}
	return config, nil
}
