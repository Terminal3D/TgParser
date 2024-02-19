-- name: InsertItem :one
INSERT INTO item (name, brand, price, available, url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;




