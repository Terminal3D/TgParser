-- name: InsertItem :one
INSERT INTO item (id, name, brand, price, available, url, chat_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;




