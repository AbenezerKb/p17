-- name: AddBalance :one
INSERT INTO balance
    (client_id, amount, status)
    VALUES ($1, $2, $3)
RETURNING *;

-- name: ListAllBalance :many
SELECT * FROM balance
      LIMIT $1
      OFFSET $2;

-- name: GetClientBalance :one
SELECT * FROM balance
      WHERE client_id=$1;

-- name: UpdateClientBalance :one
UPDATE balance
    SET  amount =  $2
    WHERE client_id=$1
    RETURNING *;
