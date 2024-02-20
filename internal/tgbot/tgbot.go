package tgbot

import (
	"TgParser/internal/database"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
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
