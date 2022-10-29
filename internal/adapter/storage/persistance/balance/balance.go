package balance

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	"strconv"
)

type balanceStorage struct {
	db  *pgxpool.Pool
	dbp db.Queries
}

type Storage interface {
	AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error)
	GetAllBalances(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error)
	GetCurrentClientBalance(ctx context.Context, clientId string) (*dto.Balance, error)
	GetLastMonthClientBalance(ctx context.Context, clientId string) (*dto.Balance, error)
	UpdateClientBalance(ctx context.Context, transfer model.Transfer, balance *dto.Balance) (*dto.Balance, error)
	GetLastMonthClientTransactions(ctx context.Context, clientId string) ([]dto.ClientTransaction, error)
	LastMonthMessagesPriceAndCount(ctx context.Context, clientId string) ([]model.MessageCount, error)
	GenerateInvoice(ctx context.Context, invoice []model.ClientInvoice) error
	ListClients(ctx context.Context) ([]dto.Client, error)
}

func StorageInit(utils const_init.Utils) Storage {
	return balanceStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
}

func (c balanceStorage) AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {

	blc, err := c.dbp.AddBalance(ctx, db.AddBalanceParams{
		ClientID: balance.ClientId,
		Amount:   balance.Amount,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)

	}

	balance = &dto.Balance{
		Id:        blc.ID.String(),
		ClientId:  blc.ClientID,
		Amount:    blc.Amount,
		Status:    blc.Status,
		CreatedAt: blc.CreatedAt,
	}

	return balance, nil
}

func (b balanceStorage) GetAllBalances(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)

	blc, err := b.ListAllBalances(ctx, ListAllBalanceParams{
		Offset: resizedPage,
		Limit:  resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return blc, nil
}

func (b balanceStorage) GetCurrentClientBalance(ctx context.Context, clientId string) (*dto.Balance, error) {

	blc, err := b.dbp.GetClientBalance(ctx, clientId)

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

func (b balanceStorage) UpdateClientBalance(ctx context.Context, transfer model.Transfer, balance *dto.Balance) (*dto.Balance, error) {

	var updatedBlc db.Balance
	dbx := pgxpool.Conn{}
	err := dbx.BeginFunc(ctx, func(tx pgx.Tx) error {
		var tfr db.Transfer
		if transfer == model.TransferCREDITING {
			tfr = db.TransferCREDITING
		} else {
			tfr = db.TransferDEBITING
		}

		_, err := b.dbp.AddTransaction(ctx, db.AddTransactionParams{
			ClientID: balance.ClientId,
			Amount:   balance.Amount,
			Type:     tfr,
		})

		if err != nil {
			return err
		}
		updatedBlc, err = b.dbp.UpdateBalance(ctx, db.UpdateBalanceParams{
			ClientID: balance.ClientId,
			Amount:   balance.Amount,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	updatedBalance := &dto.Balance{
		Id:        updatedBlc.ID.String(),
		ClientId:  updatedBlc.ClientID,
		Amount:    updatedBlc.Amount,
		Status:    updatedBlc.Status,
		CreatedAt: updatedBlc.CreatedAt,
		UpdatedAt: updatedBlc.UpdatedAt,
	}

	return updatedBalance, nil

}



func (b balanceStorage) GenerateInvoice(ctx context.Context, invoices []model.ClientInvoice) error {
	for _, invoice := range invoices {

		invoice := dto.Invoice{
			ClientId:       client.Id,
			CurrentBalance: currentBalance.Amount,
			PaymentType:    client.PaymentType,
			BalanceAtBeginning: lastMonthBalance.Amount,
			MessageCount:
		}
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
func (b balanceStorage) ListAllBalances(ctx context.Context, arg ListAllBalanceParams) ([]dto.Balance, error) {
	rows, err := b.db.Query(ctx, listAllBalance, arg.Limit, arg.Offset)
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
func (b balanceStorage) GetLastMonthClientBalance(ctx context.Context, clientId string) (*dto.Balance, error) {
	balance, err := b.dbp.GetLastMonthBalance(ctx, clientId)
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

func (b balanceStorage) GetLastMonthClientTransactions(ctx context.Context, clientId string) ([]dto.ClientTransaction, error) {
	txn, err := b.GetLastMonthClientTransaction(ctx, GetLastMonthClientTransactionParams{ClientID: clientId, Type: db.TransferCREDITING})
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

//GetLastMonthClientTransaction It list the previous month transactions for the specific client and the type of transaction
func (b balanceStorage) GetLastMonthClientTransaction(ctx context.Context, arg GetLastMonthClientTransactionParams) ([]dto.ClientTransaction, error) {
	rows, err := b.db.Query(ctx, getLastMonthClientTransaction, arg.ClientID, arg.Type)
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


func (b balanceStorage) LastMonthMessagesPriceAndCount(ctx context.Context, clientId string) ([]model.MessageCount, error){

	messageCount, err := b.LastMonthMessagePriceAndCount (ctx,clientId)
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
func (b balanceStorage) LastMonthMessagePriceAndCount(ctx context.Context, clientID string) ([]model.MessageCount, error) {
	rows, err := b.db.Query(ctx, lastMonthMessagePriceAndCount, clientID)
	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	defer rows.Close()
	items := []model.MessageCount{}
	for rows.Next() {
		var i model.MessageCount
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
func (b balanceStorage) ListClients(ctx context.Context) ([]dto.Client, error) {
	rows, err := b.db.Query(ctx, listClients)
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
