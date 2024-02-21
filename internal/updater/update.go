package updater

import (
	"TgParser/internal/data"
	"TgParser/internal/database"
	"TgParser/internal/tgbot"
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func Launch(apiDB *database.Queries, ctxIn context.Context, updateInfoChan chan data.UpdateInfo) {
	c := cron.New()
	_, err := c.AddFunc("@every 1m", func() {
		log.Println("Update cycle started")
		updateItems(apiDB, ctxIn, updateInfoChan)
	})
	if err != nil {
		log.Fatalf("Error scheduling the task: %s", err)
	}

	c.Start()
}

func updateItems(apiDB *database.Queries, ctxIn context.Context, updateInfoChan chan data.UpdateInfo) {
	// Запрос всех доступных предметов в базе данных
	items, err := apiDB.GetAllItemsWithoutSizes(ctxIn)
	if err != nil {
		log.Printf("Error fetching items: %v", err)
		return
	}

	totalItems := len(items)
	if totalItems == 0 {
		return
	}

	visited := make(map[int]bool) // Отслеживание обработанных предметов

	for len(visited) < totalItems {
		index := rand.Intn(totalItems)
		if _, ok := visited[index]; ok {
			continue
		}
		visited[index] = true
		item, err := processItem(apiDB, ctxIn, &items[index])

		if err != nil {
			log.Println(err)
			/* TODO Добавить логирование в базу данных */
		}

		updateInfoChan <- item

		waitTime := calculateInterval(totalItems)
		time.Sleep(time.Duration(waitTime) * time.Second)
	}
}

// calculateInterval вычисляет время ожидания в секундах на основе общего количества предметов
func calculateInterval(totalItems int) int {

	baseInterval := totalItems
	randomAdjustment := rand.Intn(baseInterval)
	return baseInterval + randomAdjustment
}

func processItem(apiDB *database.Queries, ctx context.Context, item *database.Item) (data.UpdateInfo, error) {

	parsedData, err := tgbot.HandleAddItem(item.Url, apiDB, ctx)
	if err != nil {
		return data.UpdateInfo{}, err
	}

	price, err := strconv.ParseFloat(item.Price, 64)
	if err != nil {
		return data.UpdateInfo{}, err
	}

	updateInfo := data.UpdateInfo{
		Status: "",
		Item:   parsedData,
	}
	switch {
	case parsedData.Available == false:
		updateInfo.Status = data.UpdateStatusNotAvailable
	case parsedData.Price > price:
		updateInfo.Status = data.UpdateStatusHigherPrice
	case parsedData.Price < price:
		updateInfo.Status = data.UpdateStatusLowerPrice
	}

	return updateInfo, nil
}
