package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5433")
	user := getEnv("DB_USER", "fleetpulse_user")
	password := getEnv("DB_PASSWORD", "fleetpulse_pass")
	dbname := getEnv("DB_NAME", "fleetpulse")

dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
    host, user, password, dbname, port,
)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully")

	return db
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
