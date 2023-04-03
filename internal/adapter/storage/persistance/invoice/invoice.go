package invoice

import (
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest/error_types"
)

type invoiceStorage struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

//Storage interface for invoice storage
type Storage interface {
	AddInvoice(ctx context.Context, clientInvoices []dto.ClientInvoice) error
	ListAllClientInvoices(ctx context.Context, clientID string) ([]dto.ClientInvoice, error)
	ListClients(ctx context.Context) ([]dto.Client, error)
	ListAllBalances(ctx context.Context, arg ListAllBalanceParams) ([]dto.Balance, error)
	GetCurrentClientBalance(ctx context.Context, clientId string) (*dto.Balance, error)
	GetLastMonthClientBalance(ctx context.Context, clientId string) (*dto.Balance, error)
	GetLastMonthClientTransactions(ctx context.Context, clientId string) ([]dto.ClientTransaction, error)
	LastMonthMessagesPriceAndCount(ctx context.Context, senderPhone string) ([]dto.MessageCount, error)
}

func StorageInit(utils const_init.Utils) Storage {
	return invoiceStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
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

//ListAllBalances lists all balances
func (is invoiceStorage) ListAllBalances(ctx context.Context, arg ListAllBalanceParams) ([]dto.Balance, error) {
	rows, err := is.db.Query(ctx, listAllBalance, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.Balance{}
	for rows.Next() {
		var i dto.Balance
		if err := rows.Scan(
			&i.Id,
			&i.ClientId,
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

//GetLastMonthClientBalance get the client balance during thr beginning previous month
func (is invoiceStorage) GetLastMonthClientBalance(ctx context.Context, clientId string) (*dto.Balance, error) {
	balance, err := is.dbp.GetLastMonthBalance(ctx, clientId)
	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	blc := dto.Balance{
		Id:        balance.ID.String(),
		ClientId:  balance.ClientID,
		Amount:    balance.Amount,
		Status:    balance.Status,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}

	return &blc, nil

}

//GetLastMonthClientTransactions It list the previous month transactions for the specific client and the type of transaction
func (is invoiceStorage) GetLastMonthClientTransactions(ctx context.Context, clientId string) ([]dto.ClientTransaction, error) {
	txn, err := is.GetLastMonthClientTransaction(ctx, GetLastMonthClientTransactionParams{ClientID: clientId, Type: db.TransferCREDITING})
	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	return txn, nil
}

const getLastMonthClientTransaction = `-- name: GetLastMonthClientTransaction :many
SELECT id, client_id, amount, type, created_at FROM client_transaction
WHERE client_id=$1 AND type=$2 AND "created_at" BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW()
`

type GetLastMonthClientTransactionParams struct {
	ClientID string      `json:"client_id"`
	Type     db.Transfer `json:"type"`
}

func (is invoiceStorage) GetLastMonthClientTransaction(ctx context.Context, arg GetLastMonthClientTransactionParams) ([]dto.ClientTransaction, error) {
	rows, err := is.db.Query(ctx, getLastMonthClientTransaction, arg.ClientID, arg.Type)
	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	defer rows.Close()
	items := []dto.ClientTransaction{}
	for rows.Next() {
		var i dto.ClientTransaction
		if err := rows.Scan(
			&i.Id,
			&i.ClientId,
			&i.Amount,
			&i.Type,
			&i.CreatedAt,
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

//LastMonthMessagesPriceAndCount searches for the previous month sent messages count and their price
func (is invoiceStorage) LastMonthMessagesPriceAndCount(ctx context.Context, clientId string) ([]dto.MessageCount, error) {

	messageCount, err := is.LastMonthMessagePriceAndCount(ctx, clientId)
	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return messageCount, err
}

const lastMonthMessagePriceAndCount = `-- name: LastMonthMessagePriceAndCount :many
SELECT  price, COUNT(id) as COUNT,
        SUM (price) AS sum
FROM public.messages
WHERE client_id=$1 AND "created_at" BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW()
GROUP BY  price
`

type LastMonthMessagePriceAndCountRow struct {
	Price decimal.Decimal `json:"price"`
	Count int64           `json:"count"`
	Sum   int64           `json:"sum"`
}

//LastMonthMessagePriceAndCount lists the previous month messages count with their sum price
func (is invoiceStorage) LastMonthMessagePriceAndCount(ctx context.Context, clientID string) ([]dto.MessageCount, error) {
	rows, err := is.db.Query(ctx, lastMonthMessagePriceAndCount, clientID)
	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	defer rows.Close()
	items := []dto.MessageCount{}
	for rows.Next() {
		var i dto.MessageCount
		if err := rows.Scan(&i.Price, &i.Count, &i.Sum); err != nil {
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

//ListClients lists all clients without pagination
func (is invoiceStorage) ListClients(ctx context.Context) ([]dto.Client, error) {
	rows, err := is.db.Query(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []dto.Client{}
	for rows.Next() {
		var i dto.Client
		if err := rows.Scan(
			&i.Id,
			&i.ClientTitle,
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

//AddInvoice insert generated invoices into database for all clients monthly
func (is invoiceStorage) AddInvoice(ctx context.Context, clientInvoices []dto.ClientInvoice) error {

	for _, invoice := range clientInvoices {
		uid, _ := uuid.Parse(invoice.InvoiceNumber)
		msgCount, err := json.Marshal(invoice.MessageCount)
		if err != nil {
			return err
		}

		clientTxn, err := json.Marshal(invoice.ClientTransactions)
		if err != nil {
			return err
		}

		var paymentType db.PaymentType
		switch invoice.PaymentType {
		case dto.PaymentTypePrepaid:
			paymentType = db.PaymentTypePrepaid
		case dto.PaymentTypePostpaid:
			paymentType = db.PaymentTypePostpaid

		}

		_, err = is.dbp.AddInvoice(ctx, db.AddInvoiceParams{
			InvoiceNumber:      uid,
			ClientID:           invoice.Id,
			PaymentType:        paymentType,
			CurrentBalance:     invoice.CurrentBalance,
			BalanceAtBeginning: invoice.BalanceAtMonthBeginning,
			MessageCount: pgtype.JSON{
				Bytes:  msgCount,
				Status: pgtype.Present,
			},
			ClientTransaction: pgtype.JSON{
				Bytes:  clientTxn,
				Status: pgtype.Present,
			},
			Tax:     invoice.Tax,
			TaxRate: invoice.TaxRate,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (is invoiceStorage) GetCurrentClientBalance(ctx context.Context, clientId string) (*dto.Balance, error) {

	blc, err := is.dbp.GetClientBalance(ctx, clientId)

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	balance := &dto.Balance{
		Id:        blc.ID.String(),
		ClientId:  blc.ClientID,
		Amount:    blc.Amount,
		Status:    blc.Status,
		CreatedAt: blc.CreatedAt,
		UpdatedAt: blc.UpdatedAt,
	}

	return balance, nil
}

const listAllClientInvoices = `-- name: ListAllClientInvoices :many
SELECT id, invoice_number, client_id, payment_type, current_balance, balance_at_beginning, message_count, client_transaction, tax, tax_rate, created_at FROM invoice
WHERE client_id=$1
`

func (is invoiceStorage) ListAllClientInvoices(ctx context.Context, clientID string) ([]dto.ClientInvoice, error) {
	rows, err := is.db.Query(ctx, listAllClientInvoices, clientID)
	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	defer rows.Close()
	items := []dto.ClientInvoice{}
	for rows.Next() {
		var i dto.ClientInvoice
		if err := rows.Scan(
			&i.Id,
			&i.InvoiceNumber,
			&i.ClientId,
			&i.PaymentType,
			&i.CurrentBalance,
			&i.BalanceAtMonthBeginning,
			&i.MessageCount,
			&i.ClientTransactions,
			&i.Tax,
			&i.TaxRate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, error_types.GetDbError(err)
	}
	return items, nil
}
