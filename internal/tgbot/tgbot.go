package tgbot

import (
	"TgParser/internal/database"
	"TgParser/internal/marketParser/models"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

var (
	startMenu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Добавить вещь",
				"add_item",
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Посмотреть список вещей",
				"list_items",
			),
		),
	)

	userStates = NewUserStates()
	apiDB      *database.Queries
	ctx        context.Context
)

func Launch(apiToken string, queries *database.Queries, ctxIn context.Context) error {
	apiDB = queries
	ctx = ctxIn
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		return err
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			if err = handleCallbacks(bot, &update); err != nil {
				log.Println(err)
			}
		} else if update.Message.IsCommand() {
			if err = handleCommand(bot, &update); err != nil {
				log.Println(err)
			}
		} else {
			if err = handleMessage(bot, &update); err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func handleCallbacks(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	data := update.CallbackQuery.Data
	chatID := update.CallbackQuery.From.ID

	if state, ok := userStates.Get(chatID); ok && state == AwaitingFilterModeState {
		err := handleFiltersPick(bot, chatID, data)
		if err != nil {
			return err
		}
	}

	switch data {
	case "add_item":
		msg := tgbotapi.NewMessage(chatID, "Укажите URL товара (введите \"отмена\", чтобы отменить запрос):")
		if _, err := bot.Send(msg); err != nil {
			return fmt.Errorf("failed to send add_item response. error: %v", msg)
		}
		userStates.Set(chatID, AwaitingUrlState)

	case "list_items":
		msg := tgbotapi.NewMessage(chatID, "Выберите тип фильтрации:")
		msg.ReplyMarkup = filterMenu
		msg.ParseMode = "Markdown"
		if _, err := bot.Send(msg); err != nil {
			return fmt.Errorf("failed to send add_item response. error: %v", msg)
		}
		userStates.Set(chatID, AwaitingFilterModeState)
	}

	return nil
}

func handleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	command := update.Message.Command()
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие")
		msg.ReplyMarkup = startMenu
		msg.ParseMode = "Markdown"

		if _, err := bot.Send(msg); err != nil {
			return fmt.Errorf("error sending start msg, error: %v", err)
		}
	}
	return nil
}

func handleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	chatID := update.Message.Chat.ID
	state, ok := userStates.Get(chatID)
	if !ok {
		msg := tgbotapi.NewMessage(chatID, "Неизвестный запрос")
		if _, err := bot.Send(msg); err != nil {
			return fmt.Errorf("error sending unknown request message")
		}
	}
	switch state {
	case AwaitingUrlState:
		defer userStates.Delete(chatID) // Сбросить состояние после обработки
		url := update.Message.Text
		if strings.ToLower(url) == "отмена" {

			return nil
		}
		productData, err := handleAddItem(url, false)
		if err != nil {
			if _, err1 := bot.Send(tgbotapi.NewMessage(chatID, err.Error())); err1 != nil {
				return fmt.Errorf("failed to send message for error: %v. sending error: %v", err, err1)
			}
			return err
		}

		if err = displayData(bot, chatID, productData); err != nil {
			return err
		}

		return nil

	case AwaitingNameFilter:
		userStates.Delete(chatID)
		name := strings.TrimSpace(update.Message.Text)

		items, err := apiDB.GetItemsByName(ctx, name)
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, models.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingBrandFilter:
		userStates.Delete(chatID)
		brand := strings.TrimSpace(update.Message.Text)

		items, err := apiDB.GetItemsByBrand(ctx, brand)
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, models.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingPriceFilter:
		userStates.Delete(chatID)
		price := strings.TrimSpace(update.Message.Text)

		items, err := apiDB.GetItemsByMaxPrice(ctx, price)
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, models.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingBrandAndPriceFilter:
		userStates.Delete(chatID)
		input := strings.Split(update.Message.Text, ",")

		brand := strings.TrimSpace(input[0])
		price := strings.TrimSpace(input[1])

		items, err := apiDB.GetItemsByBrandAndMaxPrice(ctx, database.GetItemsByBrandAndMaxPriceParams{
			Brand: brand,
			Price: price,
		})

		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, models.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil
	}
	return nil
}
