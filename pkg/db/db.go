package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Anandhu4456/rms/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func Connection(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("[Invalid dsn] :", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return &gorm.DB{}, err
	}
	if err := db.AutoMigrate(&model.Profile{}); err != nil {
		return &gorm.DB{}, err
	}
	if err := db.AutoMigrate(&model.Job{}); err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}
