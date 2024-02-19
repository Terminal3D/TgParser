-- name: InsertSize :one
INSERT INTO size (id, product_id, size, quantity)
VALUES ($1, $2, $3, $4)
RETURNING *;