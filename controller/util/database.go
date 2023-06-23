package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	Err    error
)

func PostgreSQLConnect(username, password, host, databaseName, port, sslMode, timeZone string) {

	DBConn, Err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, username, password, databaseName, port, sslMode, timeZone)),
		&gorm.Config{})
}

func CreateConnection() {
	PostgreSQLConnect(
		GetEnv("POSTGRES_USERNAME"),
		GetEnv("POSTGRES_PASSWORD"),
		GetEnv("POSTGRES_HOST"),
		GetEnv("DATABASE_NAME"),
		GetEnv("POSTGRES_PORT"),
		GetEnv("POSTGRES_SSL_MODE"),
		GetEnv("POSTGRES_TIMEZONE"),
	)
	err := DBConn.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
