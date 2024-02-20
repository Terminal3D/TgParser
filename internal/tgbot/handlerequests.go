package tgbot

import (
	"TgParser/internal/database"
	"TgParser/internal/marketParser"
	"TgParser/internal/marketParser/models"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"strings"
)

func displayData(bot *tgbotapi.BotAPI, chatID int64, data ...*models.ProductData) error {

	var responseMsg strings.Builder
	if len(data) == 0 {
		responseMsg.WriteString("Товары не найдены")
	}
	for _, productData := range data {

		responseMsg.WriteString("\n\nНазвание: " + productData.Name + "\nБрэнд: " +
			productData.Brand + "\nЦена: " + fmt.Sprintf("%.2f", productData.Price))

		for _, size := range productData.Sizes {
			responseMsg.WriteString("\nРазмер: " + size.Size + ", количество: " + fmt.Sprintf("%d", size.Quantity))
		}

	}
	msg := tgbotapi.NewMessage(chatID, responseMsg.String())
	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("error sending product data message")
	}
	return nil
}

func handleAddItem(inputURL string) (*models.ProductData, error) {

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
	}

	apiDB.DeleteItem(ctx, database.DeleteItemParams{
		Name:  insertParamsItem.Name,
		Brand: insertParamsItem.Brand,
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

func handleFiltersPick(bot *tgbotapi.BotAPI, chatID int64, data string) error {
	var text string

	switch data {
	case "no_filter":
		userStates.Delete(chatID)
		items, err := apiDB.GetAllItems(ctx)
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, models.NewFromSQL(items)...); err != nil {
			return err
		}

		if err = showStartMenu(bot, chatID); err != nil {
			return err
		}

		return nil

	case "name_filter":
		userStates.Set(chatID, AwaitingNameFilter)
		text = "Введите название товара"

	case "brand_filter":
		userStates.Set(chatID, AwaitingBrandFilter)
		text = "Введите название бренда"

	case "price_filter":
		userStates.Set(chatID, AwaitingPriceFilter)
		text = "Введите цену"

	case "brand_and_price_filter":
		userStates.Set(chatID, AwaitingBrandAndPriceFilter)
		text = "Введите название бренда и цену через запятую (напр. Puma, 24.00)"

	default:
		userStates.Delete(chatID)
		if err := showFilterMenu(bot, chatID); err != nil {
			return err
		}
	}

	if _, err := bot.Send(tgbotapi.NewMessage(chatID, text)); err != nil {
		return err
	}
	return nil
}
