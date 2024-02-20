// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: deleteItem.sql

package database

import (
	"context"
)

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM item
WHERE name = $1 AND brand = $2
`

type DeleteItemParams struct {
	Name  string
	Brand string
}

func (q *Queries) DeleteItem(ctx context.Context, arg DeleteItemParams) error {
	_, err := q.db.ExecContext(ctx, deleteItem, arg.Name, arg.Brand)
	return err
}
