-- name: AddMessage :one
INSERT INTO public.messages
    (sender_phone, content, price, receiver_phone, type,status)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListAllMessages :many
SELECT * FROM public.messages
        LIMIT $1
OFFSET $2;

-- name: GetMessagesBySender :many
SELECT * FROM public.messages
WHERE sender_phone=$1
LIMIT $2
    OFFSET $3;

-- name: GetMessageWithPrefix :many
SELECT * FROM public.messages
    WHERE content LIKE $2 AND receiver_phone=$1
    LIMIT $3
    OFFSET $4;

-- name: UpdateDeliveryStatus :exec
UPDATE public.messages
SET delivery_status = $2
WHERE id = $1;


-- name: LastMonthMessagePriceAndCount :many
SELECT  price, COUNT(id) as COUNT,
        SUM (price) AS sum
FROM public.messages
WHERE client_id=$1 AND "created_at" BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW()
GROUP BY  price;

