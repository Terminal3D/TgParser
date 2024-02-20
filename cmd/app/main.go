package main

import (
	"TgParser/internal/database"
	"TgParser/internal/tgbot"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Couldn't get DB URL")
	}

	apiToken := os.Getenv("TG_API_TOKEN")
	if apiToken == "" {
		log.Fatal("Token not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := tgbot.Launch(apiToken, database.New(conn), ctx); err != nil {
		log.Fatal(err)
	}
}
