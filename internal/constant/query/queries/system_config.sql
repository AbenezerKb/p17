-- name: AddSystemConfig :one
INSERT INTO system_config
    (setting_name,setting_value)
    VALUES ($1,$2)
    RETURNING *;


-- name: ListAllSystemConfig :many
SELECT * FROM system_config;

-- name: GetSystemConfig :one
SELECT * FROM system_config
WHERE setting_name=$1;

-- name: UpdateSystemConfig :one
UPDATE system_config
SET setting_value=$2
WHERE setting_name=$1
RETURNING *;


