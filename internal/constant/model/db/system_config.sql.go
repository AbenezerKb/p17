// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: system_config.sql

package db

import (
	"context"
)

const addSystemConfig = `-- name: AddSystemConfig :one
INSERT INTO system_config
    (setting_name,setting_value)
    VALUES ($1,$2)
    RETURNING id, setting_name, setting_value, created_at, updated_at
`

type AddSystemConfigParams struct {
	SettingName  string `json:"setting_name"`
	SettingValue string `json:"setting_value"`
}

func (q *Queries) AddSystemConfig(ctx context.Context, arg AddSystemConfigParams) (SystemConfig, error) {
	row := q.db.QueryRow(ctx, addSystemConfig, arg.SettingName, arg.SettingValue)
	var i SystemConfig
	err := row.Scan(
		&i.ID,
		&i.SettingName,
		&i.SettingValue,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSystemConfig = `-- name: GetSystemConfig :one
SELECT id, setting_name, setting_value, created_at, updated_at FROM system_config
WHERE setting_name=$1
`

func (q *Queries) GetSystemConfig(ctx context.Context, settingName string) (SystemConfig, error) {
	row := q.db.QueryRow(ctx, getSystemConfig, settingName)
	var i SystemConfig
	err := row.Scan(
		&i.ID,
		&i.SettingName,
		&i.SettingValue,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllSystemConfig = `-- name: ListAllSystemConfig :many
SELECT id, setting_name, setting_value, created_at, updated_at FROM system_config
`

func (q *Queries) ListAllSystemConfig(ctx context.Context) ([]SystemConfig, error) {
	rows, err := q.db.Query(ctx, listAllSystemConfig)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SystemConfig{}
	for rows.Next() {
		var i SystemConfig
		if err := rows.Scan(
			&i.ID,
			&i.SettingName,
			&i.SettingValue,
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

const updateSystemConfig = `-- name: UpdateSystemConfig :one
UPDATE system_config
SET setting_value=$2
WHERE setting_name=$1
RETURNING id, setting_name, setting_value, created_at, updated_at
`

type UpdateSystemConfigParams struct {
	SettingName  string `json:"setting_name"`
	SettingValue string `json:"setting_value"`
}

func (q *Queries) UpdateSystemConfig(ctx context.Context, arg UpdateSystemConfigParams) (SystemConfig, error) {
	row := q.db.QueryRow(ctx, updateSystemConfig, arg.SettingName, arg.SettingValue)
	var i SystemConfig
	err := row.Scan(
		&i.ID,
		&i.SettingName,
		&i.SettingValue,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
