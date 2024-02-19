package app

import (
	"TgParser/internal/database"
	"TgParser/internal/marketParser"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func RunApp() {

	ctx := context.Background()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Couldn't get DB URL")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}
	var command string
	for command != "stop" {
		_, err := fmt.Scan(&command)
		if err != nil {
			log.Println(err)
			return
		}
		switch command {
		case "add":
			AddItem(&apiCfg, ctx, false)
		case "update":
			AddItem(&apiCfg, ctx, true)
		default:
			command = "stop"
		}
	}

}

func AddItem(apiCfg *apiConfig, ctx context.Context, update bool) {

	var inputURL string

	_, err := fmt.Scan(&inputURL)
	if inputURL == "stop" {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	parsedData, err := marketParser.RunParser(inputURL)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Received data:", *parsedData)
	itemUUID := uuid.New()
	insertParamsItem := database.InsertItemParams{
		ID:        itemUUID,
		Name:      parsedData.Name,
		Brand:     parsedData.Brand,
		Price:     fmt.Sprintf("%.2f", parsedData.Price),
		Available: parsedData.Available,
		Url:       inputURL,
	}

	if update {
		err := apiCfg.DB.DeleteItem(ctx, database.DeleteItemParams{
			Name:  insertParamsItem.Name,
			Brand: insertParamsItem.Brand,
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

	item, err := apiCfg.DB.InsertItem(ctx, insertParamsItem)
	if err != nil {
		log.Println(err)
		return
	}

	for _, size := range parsedData.Sizes {
		insertParamsSize := database.InsertSizeParams{
			ID:        uuid.New(),
			ProductID: item.ID,
			Size:      size.Size,
			Quantity:  int32(size.Quantity),
		}
		_, err := apiCfg.DB.InsertSize(ctx, insertParamsSize)
		if err != nil {
			log.Println(err)
			break
		}
	}
	log.Println("Item Successfully added")
}
