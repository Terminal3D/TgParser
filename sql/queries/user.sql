-- name: InsertUser :one
INSERT INTO bot_user (chat_id, username, subscribed)
VALUES ($1, $2, $3)
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM bot_user
WHERE chat_id = $1;


-- name: ChangeSubscription :exec
UPDATE bot_user SET subscribed = $1 WHERE chat_id = $2;

-- name: CheckUserExists :one
SELECT EXISTS(
    SELECT 1
    FROM bot_user
    WHERE chat_id = $1
);

-- name: GetSubscribedUsers :many
SELECT chat_id, username FROM bot_user WHERE subscribed = true;