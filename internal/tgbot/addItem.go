package tgbot

import (
	"TgParser/internal/data"
	"TgParser/internal/database"
	"TgParser/internal/marketParser"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func HandleAddItem(inputURL string, chatID int64, apiDB *database.Queries, ctx context.Context) (*data.ProductData, error) {

	parsedData, err := marketParser.RunParser(inputURL)
	if err != nil {
		return nil, err
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
		ChatID:    chatID,
	}

	apiDB.DeleteItem(ctx, database.DeleteItemParams{
		Name:   insertParamsItem.Name,
		Brand:  insertParamsItem.Brand,
		ChatID: chatID,
	})

	item, err := apiDB.InsertItem(ctx, insertParamsItem)
	if err != nil {
		return nil, err
	}

	for _, size := range parsedData.Sizes {
		insertParamsSize := database.InsertSizeParams{
			ID:        uuid.New(),
			ProductID: item.ID,
			Size:      size.Size,
			Quantity:  int32(size.Quantity),
		}
		_, err := apiDB.InsertSize(ctx, insertParamsSize)
		if err != nil {
			log.Println(err)
			break
		}
	}
	return parsedData, nil
}
