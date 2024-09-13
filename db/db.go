package db

import (
	"building_service/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB(cfg *config.Config) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var err error
	DB, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
}
