// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: client.sql

package db

import (
	"context"
)

const addClient = `-- name: AddClient :one
INSERT INTO clients
    (title,phone,email,password)
VALUES ($1, $2, $3, $4)
    RETURNING id, title, phone, email, password, status, created_at, updated_at
`

type AddClientParams struct {
	Title    string   `json:"title"`
	Phone    []string `json:"phone"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
}

func (q *Queries) AddClient(ctx context.Context, arg AddClientParams) (Client, error) {
	row := q.db.QueryRow(ctx, addClient,
		arg.Title,
		arg.Phone,
		arg.Email,
		arg.Password,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Phone,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getClient = `-- name: GetClient :one
SELECT id, title, phone, email, password, status, created_at, updated_at FROM clients
WHERE email=$1
`

func (q *Queries) GetClient(ctx context.Context, email string) (Client, error) {
	row := q.db.QueryRow(ctx, getClient, email)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Phone,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllClients = `-- name: ListAllClients :many
SELECT id, title, phone, email, password, status, created_at, updated_at FROM clients
LIMIT $1
OFFSET $2
`

type ListAllClientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllClients(ctx context.Context, arg ListAllClientsParams) ([]Client, error) {
	rows, err := q.db.Query(ctx, listAllClients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Client{}
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Phone,
			&i.Email,
			&i.Password,
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

const listClients = `-- name: ListClients :many
SELECT id, title, phone, email, password, status, created_at, updated_at FROM clients
`

func (q *Queries) ListClients(ctx context.Context) ([]Client, error) {
	rows, err := q.db.Query(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Client{}
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Phone,
			&i.Email,
			&i.Password,
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

const updateClient = `-- name: UpdateClient :one
UPDATE clients SET
    title=$2,
    email=$3,
    password=$4,
    status=$5,
    updated_at=NOW()
    WHERE phone=$1
    RETURNING id, title, phone, email, password, status, created_at, updated_at
`

type UpdateClientParams struct {
	Phone    []string `json:"phone"`
	Title    string   `json:"title"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Status   string   `json:"status"`
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) (Client, error) {
	row := q.db.QueryRow(ctx, updateClient,
		arg.Phone,
		arg.Title,
		arg.Email,
		arg.Password,
		arg.Status,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Phone,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
