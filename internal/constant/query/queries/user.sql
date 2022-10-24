-- name: AddUser :one
INSERT INTO users
    (full_name,phone, password)
    VALUES ($1, $2, $3)
    RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
LIMIT $1
OFFSET $2;

-- name: GetUser :one
SELECT * FROM users
WHERE phone=$1;

-- name: UpdateUser :one
UPDATE users
    SET full_name=$2, updated_at=now()
    WHERE phone=$1
    RETURNING *;
