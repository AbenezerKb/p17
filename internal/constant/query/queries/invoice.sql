-- name: AddInvoice :one
INSERT INTO invoice
(invoice_number,client_id, payment, amount, total_monthly, discount, tax, tax_rate)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: ListClientInvoices :many
SELECT * FROM invoice
WHERE client_id=$1
LIMIT $2
OFFSET $3;


-- name: GetInvoices :one
SELECT * FROM invoice
WHERE invoice_number=$1;
