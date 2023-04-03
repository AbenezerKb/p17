-- name: AddTemplate :one
INSERT INTO templates
(client,template, category)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListAllTemplates :many
SELECT * FROM templates
LIMIT $1
OFFSET $2;

-- name: ListClientTemplates :many
SELECT * FROM templates
    WHERE client=$1
    LIMIT $2
    OFFSET $3;

-- name: UpdateTemplate :one
UPDATE templates
SET template=$2, category=$3, updated_at=now()
WHERE client=$1 AND template_id=$2
RETURNING *;
