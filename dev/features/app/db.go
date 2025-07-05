package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func StartDBConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable port=5432",
		os.Getenv("DBuser"),
		os.Getenv("DBpass"),
		"stock_checker",
	)

	client, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return client
}
