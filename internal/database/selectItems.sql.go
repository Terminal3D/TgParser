// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: selectItems.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getAllItems = `-- name: GetAllItems :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE chat_id = $1
`

type GetAllItemsRow struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
	Size      string
	Quantity  int32
}

func (q *Queries) GetAllItems(ctx context.Context, chatID int64) ([]GetAllItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllItems, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllItemsRow
	for rows.Next() {
		var i GetAllItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Brand,
			&i.Price,
			&i.Available,
			&i.Url,
			&i.Size,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllItemsWithoutSizes = `-- name: GetAllItemsWithoutSizes :many
SELECT id, name, brand, price, available, url, last_check, chat_id FROM item WHERE available = true
`

func (q *Queries) GetAllItemsWithoutSizes(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getAllItemsWithoutSizes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Brand,
			&i.Price,
			&i.Available,
			&i.Url,
			&i.LastCheck,
			&i.ChatID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsByBrand = `-- name: GetItemsByBrand :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE brand = $1 AND chat_id = $2
`

type GetItemsByBrandParams struct {
	Brand  string
	ChatID int64
}

type GetItemsByBrandRow struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
	Size      string
	Quantity  int32
}

func (q *Queries) GetItemsByBrand(ctx context.Context, arg GetItemsByBrandParams) ([]GetItemsByBrandRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByBrand, arg.Brand, arg.ChatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsByBrandRow
	for rows.Next() {
		var i GetItemsByBrandRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Brand,
			&i.Price,
			&i.Available,
			&i.Url,
			&i.Size,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsByBrandAndMaxPrice = `-- name: GetItemsByBrandAndMaxPrice :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE brand = $1 AND price <= $2 AND chat_id = $3
`

type GetItemsByBrandAndMaxPriceParams struct {
	Brand  string
	Price  string
	ChatID int64
}

type GetItemsByBrandAndMaxPriceRow struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
	Size      string
	Quantity  int32
}

func (q *Queries) GetItemsByBrandAndMaxPrice(ctx context.Context, arg GetItemsByBrandAndMaxPriceParams) ([]GetItemsByBrandAndMaxPriceRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByBrandAndMaxPrice, arg.Brand, arg.Price, arg.ChatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsByBrandAndMaxPriceRow
	for rows.Next() {
		var i GetItemsByBrandAndMaxPriceRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Brand,
			&i.Price,
			&i.Available,
			&i.Url,
			&i.Size,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsByMaxPrice = `-- name: GetItemsByMaxPrice :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE price <= $1 AND chat_id = $2
`

type GetItemsByMaxPriceParams struct {
	Price  string
	ChatID int64
}

type GetItemsByMaxPriceRow struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
	Size      string
	Quantity  int32
}

func (q *Queries) GetItemsByMaxPrice(ctx context.Context, arg GetItemsByMaxPriceParams) ([]GetItemsByMaxPriceRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByMaxPrice, arg.Price, arg.ChatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsByMaxPriceRow
	for rows.Next() {
		var i GetItemsByMaxPriceRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Brand,
			&i.Price,
			&i.Available,
			&i.Url,
			&i.Size,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemsByName = `-- name: GetItemsByName :many
SELECT item.id, item.name, item.brand, item.price, item.available, item.url, size.size, size.quantity
FROM item JOIN size ON item.id = size.product_id
WHERE name = $1 AND chat_id = $2
`

type GetItemsByNameParams struct {
	Name   string
	ChatID int64
}

type GetItemsByNameRow struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
	Size      string
	Quantity  int32
}

func (q *Queries) GetItemsByName(ctx context.Context, arg GetItemsByNameParams) ([]GetItemsByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getItemsByName, arg.Name, arg.ChatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsByNameRow
	for rows.Next() {
		var i GetItemsByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Brand,
			&i.Price,
			&i.Available,
			&i.Url,
			&i.Size,
			&i.Quantity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
