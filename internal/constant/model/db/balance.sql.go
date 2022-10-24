// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: balance.sql

package db

import (
	"context"

	"github.com/shopspring/decimal"
)

const addBalance = `-- name: AddBalance :one
INSERT INTO balance
    (client_id, amount, status)
    VALUES ($1, $2, $3)
RETURNING id, client_id, amount, status, created_at, updated_at
`

type AddBalanceParams struct {
	ClientID string          `json:"client_id"`
	Amount   decimal.Decimal `json:"amount"`
	Status   string          `json:"status"`
}

func (q *Queries) AddBalance(ctx context.Context, arg AddBalanceParams) (Balance, error) {
	row := q.db.QueryRow(ctx, addBalance, arg.ClientID, arg.Amount, arg.Status)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getClientBalance = `-- name: GetClientBalance :one
SELECT id, client_id, amount, status, created_at, updated_at FROM balance
      WHERE client_id=$1
`

func (q *Queries) GetClientBalance(ctx context.Context, clientID string) (Balance, error) {
	row := q.db.QueryRow(ctx, getClientBalance, clientID)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllBalance = `-- name: ListAllBalance :many
SELECT id, client_id, amount, status, created_at, updated_at FROM balance
      LIMIT $1
      OFFSET $2
`

type ListAllBalanceParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllBalance(ctx context.Context, arg ListAllBalanceParams) ([]Balance, error) {
	rows, err := q.db.Query(ctx, listAllBalance, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Balance{}
	for rows.Next() {
		var i Balance
		if err := rows.Scan(
			&i.ID,
			&i.ClientID,
			&i.Amount,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClientBalance = `-- name: UpdateClientBalance :one
UPDATE balance
    SET  amount =  $2
    WHERE client_id=$1
    RETURNING id, client_id, amount, status, created_at, updated_at
`

type UpdateClientBalanceParams struct {
	ClientID string          `json:"client_id"`
	Amount   decimal.Decimal `json:"amount"`
}

func (q *Queries) UpdateClientBalance(ctx context.Context, arg UpdateClientBalanceParams) (Balance, error) {
	row := q.db.QueryRow(ctx, updateClientBalance, arg.ClientID, arg.Amount)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}