package config

import (
	"log"
)

// Config структура для хранения конфигурации базы данных
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig функция для загрузки конфигурации
func LoadConfig() *Config {
	config := &Config{
		DBHost:     "db",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "12345678",
		DBName:     "building_service",
	}
	log.Println("DB config loaded successfully")
	return config
}
