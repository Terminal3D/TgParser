package tgbot

import (
	"TgParser/internal/data"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func displayData(bot *tgbotapi.BotAPI, chatID int64, data ...*data.ProductData) error {

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

func handleFiltersPick(bot *tgbotapi.BotAPI, chatID int64, callbackData string) error {
	var text string

	switch callbackData {
	case "no_filter":
		userStates.Delete(chatID)
		items, err := apiDB.GetAllItems(ctx)
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, data.NewFromSQL(items)...); err != nil {
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
		if err := showStartMenu(bot, chatID); err != nil {
			return err
		}
		return nil
	}

	if _, err := bot.Send(tgbotapi.NewMessage(chatID, text)); err != nil {
		return err
	}
	return nil
}
