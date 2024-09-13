package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		DBHost:     os.Getenv("HOST"),
		DBPort:     os.Getenv("PORT"),
		DBUser:     os.Getenv("USER"),
		DBPassword: os.Getenv("PASSWORD"),
		DBName:     os.Getenv("DBNAME"),
	}
	fmt.Println("db working")
	return config, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
