package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var startMenu = tgbotapi.NewInlineKeyboardMarkup(
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
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Удалить вещь",
			"delete_item",
		),
	),
)

var filterMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"Без фильтров",
			"no_filter",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"По названию",
			"name_filter",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"По бренду",
			"brand_filter",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"По цене",
			"price_filter",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"По бренду и цене",
			"brand_and_price_filter",
		),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(
			"отмена",
			"back",
		),
	),
)

func showFilterMenu(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите тип фильтрации")
	msg.ReplyMarkup = filterMenu
	msg.ParseMode = "Markdown"

	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("error sending filter msg, error: %v", err)
	}
	userStates.Set(chatID, AwaitingFilterModeState)
	return nil
}

func showStartMenu(bot *tgbotapi.BotAPI, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Выберите действие")
	msg.ReplyMarkup = startMenu
	msg.ParseMode = "Markdown"

	if _, err := bot.Send(msg); err != nil {
		return fmt.Errorf("error sending start msg, error: %v", err)
	}
	return nil
}
