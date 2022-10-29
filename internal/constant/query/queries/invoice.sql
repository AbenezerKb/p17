-- name: AddInvoice :one
INSERT INTO invoice
(invoice_number,client_id, payment_type, current_balance, balance_at_beginning, message_count, client_transaction, discount, tax, tax_rate)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;


-- name: ListClientInvoices :many
SELECT * FROM invoice
WHERE client_id=$1
LIMIT $2
OFFSET $3;


-- name: GetInvoice :one
SELECT * FROM invoice
WHERE invoice_number=$1;

-- name: GetInvoiceByMonth :one
SELECT * FROM invoice
WHERE client_id=$1
AND created_at=$2;
