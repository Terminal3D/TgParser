// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type BotUser struct {
	ID         int32
	ChatID     int64
	Username   sql.NullString
	Subscribed bool
}

type Item struct {
	ID        uuid.UUID
	Name      string
	Brand     string
	Price     string
	Available bool
	Url       string
	LastCheck sql.NullTime
}

type Size struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	Size      string
	Quantity  int32
}
