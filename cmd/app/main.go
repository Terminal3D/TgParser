package main

import (
	"TgParser/internal/data"
	"TgParser/internal/database"
	"TgParser/internal/tgbot"
	"TgParser/internal/updater"
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

	apiDB := database.New(conn)

	updateInfoChan := make(chan data.UpdateInfo, 4)

	go updater.Launch(apiDB, ctx, updateInfoChan)

	if err := tgbot.Launch(apiToken, apiDB, ctx, updateInfoChan); err != nil {
		log.Fatal(err)
	}

}
