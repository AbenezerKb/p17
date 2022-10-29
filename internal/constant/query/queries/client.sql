-- name: AddClient :one
INSERT INTO clients
    (title,phone,email,password)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: ListAllClients :many
SELECT * FROM clients
LIMIT $1
OFFSET $2;

-- name: ListClients :many
SELECT * FROM clients;

-- name: GetClient :one
SELECT * FROM clients
WHERE email=$1;

-- name: UpdateClient :one
UPDATE clients SET
    title=$2,
    email=$3,
    password=$4,
    status=$5,
    updated_at=NOW()
    WHERE phone=$1
    RETURNING *;