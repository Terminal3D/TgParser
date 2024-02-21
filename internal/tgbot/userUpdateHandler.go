package tgbot

import (
	"TgParser/internal/data"
	"TgParser/internal/database"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func handleCallbacks(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	callbackData := update.CallbackQuery.Data
	chatID := update.CallbackQuery.From.ID

	if state, ok := userStates.Get(chatID); ok && state == AwaitingFilterModeState {
		err := handleFiltersPick(bot, chatID, callbackData)
		if err != nil {
			return err
		}
	}

	switch callbackData {
	case "add_item":
		msg := tgbotapi.NewMessage(chatID, "Укажите URL товара (введите \"отмена\", чтобы отменить запрос):")
		if _, err := bot.Send(msg); err != nil {
			return fmt.Errorf("failed to send add_item response. error: %v", msg)
		}
		userStates.Set(chatID, AwaitingUrlState)

	case "list_items":
		if err := showFilterMenu(bot, chatID); err != nil {
			return err
		}

	case "delete_item":
		msg := tgbotapi.NewMessage(chatID, "Укажите бренд и товар через запятую (напр. Puma, T-Shirt)")
		if _, err := bot.Send(msg); err != nil {
			return fmt.Errorf("failed to send delete_item response. error: %v", msg)
		}
		userStates.Set(chatID, AwaitingDeleteItem)
	}

	return nil
}

func handleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	command := update.Message.Command()
	chatID := update.Message.Chat.ID
	username := update.Message.Chat.UserName

	subscribeMsg := "Вы успешно подписались на обновления о товарах. " +
		"Введите /unsubscribe чтобы отписаться."

	switch command {
	case "start":
		status, err := apiDB.CheckUserExists(ctx, chatID)
		if err != nil {
			return err
		}

		if !status {
			_, err := apiDB.InsertUser(ctx, database.InsertUserParams{
				ChatID:     chatID,
				Username:   sql.NullString{String: username, Valid: true},
				Subscribed: true,
			})

			if err != nil {
				log.Printf("Ошибка при добавлении пользователя: %v", err)
				return err
			}
			if _, err = bot.Send(tgbotapi.NewMessage(chatID, subscribeMsg)); err != nil {
				return err
			}
		}

		if err := showStartMenu(bot, chatID); err != nil {
			return err
		}

	case "stop":
		err := apiDB.DeleteUser(ctx, chatID)
		if err != nil {
			log.Printf("Ошибка при удалении пользователя: %v", err)
			return err
		}
		if _, err = bot.Send(tgbotapi.NewMessage(chatID, "Все данные успешно удалены.")); err != nil {
			return err
		}

	case "subscribe":
		err := apiDB.ChangeSubscription(ctx, database.ChangeSubscriptionParams{
			Subscribed: true,
			ChatID:     chatID,
		})
		if err != nil {
			log.Printf("Ошибка при подписке пользователя: %v", err)
			return err
		}
		if _, err = bot.Send(tgbotapi.NewMessage(chatID, "Вы подписались на обновления.")); err != nil {
			return err
		}

	case "unsubscribe":
		err := apiDB.ChangeSubscription(ctx, database.ChangeSubscriptionParams{
			Subscribed: false,
			ChatID:     chatID,
		})
		if err != nil {
			log.Printf("Ошибка при отписке пользователя: %v", err)
			return err
		}
		if _, err = bot.Send(tgbotapi.NewMessage(chatID, "Вы отписались от обновлений.")); err != nil {
			return err
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
		productData, err := HandleAddItem(url, chatID, apiDB, ctx)
		if err != nil {
			if _, err1 := bot.Send(tgbotapi.NewMessage(chatID, err.Error())); err1 != nil {
				return fmt.Errorf("failed to send message for error: %v. sending error: %v", err, err1)
			}
			return err
		}
		if productData.Available {
			if err = displayData(bot, chatID, productData); err != nil {
				return err
			}
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, "Предмет не доступен (некорректный URL или нет в наличии)"))
		}

		if err = showStartMenu(bot, update.Message.Chat.ID); err != nil {
			return err
		}
		return nil

	case AwaitingNameFilter:
		userStates.Delete(chatID)
		name := strings.TrimSpace(update.Message.Text)

		items, err := apiDB.GetItemsByName(ctx, database.GetItemsByNameParams{
			Name:   name,
			ChatID: chatID,
		})
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, data.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingBrandFilter:
		userStates.Delete(chatID)
		brand := strings.TrimSpace(update.Message.Text)

		items, err := apiDB.GetItemsByBrand(ctx, database.GetItemsByBrandParams{
			Brand:  brand,
			ChatID: chatID,
		})
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, data.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingPriceFilter:
		userStates.Delete(chatID)
		price := strings.TrimSpace(update.Message.Text)

		items, err := apiDB.GetItemsByMaxPrice(ctx, database.GetItemsByMaxPriceParams{
			Price:  price,
			ChatID: chatID,
		})
		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, data.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingBrandAndPriceFilter:
		userStates.Delete(chatID)
		input := strings.Split(update.Message.Text, ",")
		if len(input) != 2 {
			if _, err := bot.Send(tgbotapi.NewMessage(chatID, "Неправильный формат ввода.")); err != nil {
				return err
			}

			if err := showFilterMenu(bot, chatID); err != nil {
				return err
			}

			return fmt.Errorf("wrong input format")
		}
		brand := strings.TrimSpace(input[0])
		price := strings.TrimSpace(input[1])

		items, err := apiDB.GetItemsByBrandAndMaxPrice(ctx, database.GetItemsByBrandAndMaxPriceParams{
			Brand: brand,
			Price: price,
		})

		if err != nil {
			return err
		}

		if err = displayData(bot, chatID, data.NewFromSQL(items)...); err != nil {
			return err
		}
		return nil

	case AwaitingDeleteItem:
		userStates.Delete(chatID)
		input := strings.Split(update.Message.Text, ",")
		if len(input) != 2 {
			if _, err := bot.Send(tgbotapi.NewMessage(chatID, "Неправильный формат ввода. Пример: Puma, T-Shirt")); err != nil {
				userStates.Set(chatID, AwaitingDeleteItem)
				return err
			}

			return fmt.Errorf("wrong delete format")
		}
		brand := strings.TrimSpace(input[0])
		name := strings.TrimSpace(input[1])
		if err := apiDB.DeleteItem(ctx, database.DeleteItemParams{
			Name:  name,
			Brand: brand,
		}); err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка удаления: %v", err)))
			return err
		}
		_, err := bot.Send(tgbotapi.NewMessage(chatID, "Предмет успешно удалён"))
		if err != nil {
			return err
		}

	}
	return nil
}
