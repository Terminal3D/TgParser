package app

import (
	"TgParser/internal/database"
	"TgParser/internal/marketParser"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func RunApp() {
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

	for err == nil {
		var inputURL string
		_, err := fmt.Scan(&inputURL)

		if err != nil {
			log.Println(err)
			break
		}
		parsedData, err := marketParser.RunParser(inputURL)
		if err != nil {
			log.Println(err)
			break
		}

		fmt.Println("Received data:", *parsedData)
		insertParamsItem := database.InsertItemParams{
			Name:      parsedData.Name,
			Brand:     parsedData.Brand,
			Price:     fmt.Sprintf("%.2f", parsedData.Price),
			Available: parsedData.Available,
			Url:       inputURL,
		}

		item, err := apiCfg.DB.InsertItem(context.Background(), insertParamsItem)
		if err != nil {
			log.Println(err)
			break
		}

		for _, size := range parsedData.Sizes {
			insertParamsSize := database.InsertSizeParams{
				ProductID: item.ID,
				Size:      size.Size,
				Quantity:  int32(size.Quantity),
			}
			_, err := apiCfg.DB.InsertSize(context.Background(), insertParamsSize)
			if err != nil {
				log.Println(err)
				break
			}
		}
		log.Println("Item Successfully added")
	}
}
