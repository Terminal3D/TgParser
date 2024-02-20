package models

import (
	"TgParser/internal/database"
	"github.com/google/uuid"
	"strconv"
)

type ProductData struct {
	Name      string     `json:"name"`
	Brand     string     `json:"brand"`
	Available bool       `json:"available"`
	Price     float64    `json:"price"`
	URL       string     `json:"URL"`
	Sizes     []SizeData `json:"sizes"`
}

type SizeData struct {
	Size     string `json:"size"`
	Quantity int    `json:"qty"`
}

type DBAnswer interface {
	GetID() uuid.UUID
	GetName() string
	GetBrand() string
	GetPrice() string
	GetAvailable() bool
	GetUrl() string
	GetSize() string
	GetQuantity() int32
}

func NewFromSQL[T any](items []T) []*ProductData {
	productsResponse := make([]*ProductData, 0)

	if len(items) == 0 {
		return nil
	}

	var currentProduct *ProductData
	var currentID uuid.UUID
	var currentSizes []SizeData

	newProduct := func(dbItem DBAnswer) {
		currentProduct.Sizes = currentSizes
		productsResponse = append(productsResponse, currentProduct)

		currentSizes = []SizeData{}
	}
	var dbItem DBAnswer
	for _, item := range items {

		switch v := any(item).(type) {
		case database.GetAllItemsRow:
			dbItem = AllItemsRowWrapper{Item: &v}
		case database.GetItemsByNameRow:
			dbItem = ItemsByNameRowWrapper{Item: &v}
		case database.GetItemsByMaxPriceRow:
			dbItem = ItemsByMaxPriceRowWrapper{Item: &v}
		case database.GetItemsByBrandRow:
			dbItem = ItemsByBrandRowWrapper{Item: &v}
		case database.GetItemsByBrandAndMaxPriceRow:
			dbItem = ItemsByBrandAndPriceRowWrapper{Item: &v}
		default:
			continue
		}

		if currentProduct != nil && dbItem.GetID() != currentID {
			// Сохраняем предыдущий продукт и начинаем новый
			newProduct(dbItem)
		}

		if currentProduct == nil || dbItem.GetID() != currentID {
			// Начало нового продукта
			price, _ := strconv.ParseFloat(dbItem.GetPrice(), 64)
			currentProduct = &ProductData{
				Name:      dbItem.GetName(),
				Brand:     dbItem.GetBrand(),
				Available: dbItem.GetAvailable(),
				URL:       dbItem.GetUrl(),
				Price:     price,
				Sizes:     []SizeData{},
			}
			currentID = dbItem.GetID()
		}

		currentSizes = append(currentSizes, SizeData{
			Size:     dbItem.GetSize(),
			Quantity: int(dbItem.GetQuantity()),
		})
	}

	newProduct(dbItem)

	return productsResponse
}
