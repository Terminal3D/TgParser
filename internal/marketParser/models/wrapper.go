package models

import (
	"TgParser/internal/database"
	"github.com/google/uuid"
)

type AllItemsRowWrapper struct {
	Item *database.GetAllItemsRow
}

func (iw AllItemsRowWrapper) GetID() uuid.UUID {
	return iw.Item.ID
}

func (iw AllItemsRowWrapper) GetName() string {
	return iw.Item.Name
}

func (iw AllItemsRowWrapper) GetBrand() string {
	return iw.Item.Brand
}

func (iw AllItemsRowWrapper) GetPrice() string {
	return iw.Item.Price
}

func (iw AllItemsRowWrapper) GetAvailable() bool {
	return iw.Item.Available
}

func (iw AllItemsRowWrapper) GetUrl() string {
	return iw.Item.Url
}

func (iw AllItemsRowWrapper) GetQuantity() int32 {
	return iw.Item.Quantity
}

func (iw AllItemsRowWrapper) GetSize() string {
	return iw.Item.Size
}

type ItemsByNameRowWrapper struct {
	Item *database.GetItemsByNameRow
}

func (iw ItemsByNameRowWrapper) GetID() uuid.UUID {
	return iw.Item.ID
}

func (iw ItemsByNameRowWrapper) GetName() string {
	return iw.Item.Name
}

func (iw ItemsByNameRowWrapper) GetBrand() string {
	return iw.Item.Brand
}

func (iw ItemsByNameRowWrapper) GetPrice() string {
	return iw.Item.Price
}

func (iw ItemsByNameRowWrapper) GetAvailable() bool {
	return iw.Item.Available
}

func (iw ItemsByNameRowWrapper) GetUrl() string {
	return iw.Item.Url
}

func (iw ItemsByNameRowWrapper) GetQuantity() int32 {
	return iw.Item.Quantity
}

func (iw ItemsByNameRowWrapper) GetSize() string {
	return iw.Item.Size
}

type ItemsByMaxPriceRowWrapper struct {
	Item *database.GetItemsByMaxPriceRow
}

func (iw ItemsByMaxPriceRowWrapper) GetID() uuid.UUID {
	return iw.Item.ID
}

func (iw ItemsByMaxPriceRowWrapper) GetName() string {
	return iw.Item.Name
}

func (iw ItemsByMaxPriceRowWrapper) GetBrand() string {
	return iw.Item.Brand
}

func (iw ItemsByMaxPriceRowWrapper) GetPrice() string {
	return iw.Item.Price
}

func (iw ItemsByMaxPriceRowWrapper) GetAvailable() bool {
	return iw.Item.Available
}

func (iw ItemsByMaxPriceRowWrapper) GetUrl() string {
	return iw.Item.Url
}

func (iw ItemsByMaxPriceRowWrapper) GetQuantity() int32 {
	return iw.Item.Quantity
}

func (iw ItemsByMaxPriceRowWrapper) GetSize() string {
	return iw.Item.Size
}

type ItemsByBrandAndPriceRowWrapper struct {
	Item *database.GetItemsByBrandAndMaxPriceRow
}

func (iw ItemsByBrandAndPriceRowWrapper) GetID() uuid.UUID {
	return iw.Item.ID
}

func (iw ItemsByBrandAndPriceRowWrapper) GetName() string {
	return iw.Item.Name
}

func (iw ItemsByBrandAndPriceRowWrapper) GetBrand() string {
	return iw.Item.Brand
}

func (iw ItemsByBrandAndPriceRowWrapper) GetPrice() string {
	return iw.Item.Price
}

func (iw ItemsByBrandAndPriceRowWrapper) GetAvailable() bool {
	return iw.Item.Available
}

func (iw ItemsByBrandAndPriceRowWrapper) GetUrl() string {
	return iw.Item.Url
}

func (iw ItemsByBrandAndPriceRowWrapper) GetQuantity() int32 {
	return iw.Item.Quantity
}

func (iw ItemsByBrandAndPriceRowWrapper) GetSize() string {
	return iw.Item.Size
}

type ItemsByBrandRowWrapper struct {
	Item *database.GetItemsByBrandRow
}

func (iw ItemsByBrandRowWrapper) GetID() uuid.UUID {
	return iw.Item.ID
}

func (iw ItemsByBrandRowWrapper) GetName() string {
	return iw.Item.Name
}

func (iw ItemsByBrandRowWrapper) GetBrand() string {
	return iw.Item.Brand
}

func (iw ItemsByBrandRowWrapper) GetPrice() string {
	return iw.Item.Price
}

func (iw ItemsByBrandRowWrapper) GetAvailable() bool {
	return iw.Item.Available
}

func (iw ItemsByBrandRowWrapper) GetUrl() string {
	return iw.Item.Url
}

func (iw ItemsByBrandRowWrapper) GetQuantity() int32 {
	return iw.Item.Quantity
}

func (iw ItemsByBrandRowWrapper) GetSize() string {
	return iw.Item.Size
}
