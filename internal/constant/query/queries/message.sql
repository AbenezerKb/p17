-- name: AddMessage :one
INSERT INTO messages
    (sender_phone, content, price, receiver_phone, type,status)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListAllMessages :many
SELECT * FROM messages
        LIMIT $1
OFFSET $2;

-- name: GetMessagesBySender :many
SELECT * FROM messages
WHERE sender_phone=$1
LIMIT $2
    OFFSET $3;

-- name: GetMessageWithPrefix :many
SELECT * FROM messages
    WHERE content LIKE $2 AND receiver_phone=$1
    LIMIT $3
    OFFSET $4;

-- name: UpdateDeliveryStatus :exec
UPDATE messages
SET delivery_status = $2
WHERE message_id = $1;