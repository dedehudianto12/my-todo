package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func InitDB()*pgxpool.Pool{
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database_url := os.Getenv("DATABASE_URL")

	conn, err := pgxpool.New(context.Background(), database_url)

	if err != nil {
		log.Fatalf("Error when connecting to database")
	}

	fmt.Println("Connected with database")

	return conn
}