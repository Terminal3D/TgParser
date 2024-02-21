package tgbot

import (
	"TgParser/internal/data"
	"TgParser/internal/database"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

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
			msg = fmt.Sprintf("Данные о товаре не изменились.\n")
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
