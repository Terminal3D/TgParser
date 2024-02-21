package tgbot

import (
	"TgParser/internal/data"
	"TgParser/internal/database"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	userStates = NewUserStates()
	apiDB      *database.Queries
	ctx        context.Context
)

func Launch(apiToken string, queries *database.Queries, ctxIn context.Context, updateInfoChan chan data.UpdateInfo) error {
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

	go handleUpdateInfo(updateInfoChan, bot)

	for update := range updates {
		go func(upd tgbotapi.Update) { // Ассинхронный запуск обработчиков
			if upd.CallbackQuery != nil {
				if err := handleCallbacks(bot, &upd); err != nil {
					log.Println(err)
				}
			} else if upd.Message.IsCommand() {
				if err := handleCommand(bot, &upd); err != nil {
					log.Println(err)
				}
			} else {
				if err := handleMessage(bot, &upd); err != nil {
					log.Println(err)
				}
			}
		}(update)
	}
	return nil
}

func handleUpdateInfo(updateInfoChan chan data.UpdateInfo, bot *tgbotapi.BotAPI) {
	for updateInfo := range updateInfoChan {
		log.Println("Получил информации об обновлении: ", *updateInfo.Item)
		var msg string
		switch updateInfo.Status {
		case data.UpdateStatusNotAvailable:
			msg = fmt.Sprintf("Товар больше не доступен. ")
		case data.UpdateStatusLowerPrice:
			msg = fmt.Sprintf("Цена на товар уменьшилась. Предыдущая: %.2f, текущая: %.2f", updateInfo.PreviousPrice, updateInfo.CurrentPrice)
		case data.UpdateStatusHigherPrice:
			msg = fmt.Sprintf("Цена на товар увеличилась. Предыдущая: %.2f, текущая: %.2f", updateInfo.PreviousPrice, updateInfo.CurrentPrice)
		default:
			msg = fmt.Sprintf("Ничего о товаре не изменилось. ")
			// continue
		}
		outputMessage := fmt.Sprintf("%sДанные о товаре: \n"+
			"Бренд: %s\nНазвание: %s\nПоследняя цена: %.2f\nURL: %s\n",
			msg, updateInfo.Item.Brand, updateInfo.Item.Name, updateInfo.Item.Price, updateInfo.Item.URL)

		users, err := apiDB.GetSubscribedUsers(ctx)
		if err != nil {
			log.Println("Failed to get users")
		}
		for _, user := range users {
			_, err := bot.Send(tgbotapi.NewMessage(user.ChatID, outputMessage))
			if err != nil {
				apiDB.ChangeSubscription(ctx, database.ChangeSubscriptionParams{
					Subscribed: false,
					ChatID:     user.ChatID,
				})
			}
		}
	}
}
