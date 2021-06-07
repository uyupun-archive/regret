package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDSN() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	parseTime := "true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s", userName, password, host, port, name, parseTime)

	return dsn, nil
}
