-- name: DeleteItem :exec
DELETE FROM item
WHERE name = $1 AND brand = $2;