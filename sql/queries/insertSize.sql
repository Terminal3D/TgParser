-- name: InsertSize :one
INSERT INTO size (product_id, size, quantity)
VALUES ($1, $2, $3)
RETURNING *;