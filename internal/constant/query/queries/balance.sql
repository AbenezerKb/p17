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

-- name: UpdateBalance :one
UPDATE balance
    SET  amount =  $2,
         updated_at = now()
    WHERE client_id=$1
    RETURNING *;

-- name: AddTransaction :one
INSERT INTO client_transaction
    (client_id, amount, type)
    VALUES ($1, $2, $3)
    RETURNING *;

-- name: GetLastMonthBalance :one
SELECT * FROM balance
WHERE client_id=$1
  AND  updated_at >= date_trunc('month', current_date - interval '1' month)
  AND updated_at < date_trunc('month', current_date)
  ORDER BY updated_at DESC
  LIMIT 1;

-- name: GetLastMonthClientTransaction :many
SELECT * FROM client_transaction
WHERE client_id=$1 AND type=$2 AND "created_at" BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW();
