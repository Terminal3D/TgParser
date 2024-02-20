// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: insertItem.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const insertItem = `-- name: InsertItem :one
INSERT INTO item (id, name, brand, price, available, url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, brand, price, available, url, last_check
`

type InsertItemParams struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, insertItem,
		arg.ID,
		arg.Name,
		arg.Brand,
		arg.Price,
		arg.Available,
		arg.Url,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Brand,
		&i.Price,
		&i.Available,
		&i.Url,
		&i.LastCheck,
	)
	return i, err
}
