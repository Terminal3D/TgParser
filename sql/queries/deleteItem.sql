-- name: DeleteItemById :exec
DELETE FROM item
WHERE id = $1;