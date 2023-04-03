package balance

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
	UpdateClientBalance(ctx context.Context, transfer model.Transfer, balance *dto.Balance) (*dto.Balance, error)
	GetCreditedTransactions(ctx context.Context, ClientID string, Type db.Transfer) ([]dto.ClientTransaction, error)
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

//GetAllBalances lists all balances
func (bs balanceStorage) GetAllBalances(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)

	blc, err := bs.ListAllBalances(ctx, ListAllBalanceParams{
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

func (b balanceStorage) GetCreditedTransactions(ctx context.Context, ClientID string, Type db.Transfer) ([]dto.ClientTransaction, error) {
	txn, err := b.GetCreditedTransaction(ctx, GetCreditedTransactionParams{
		ClientID,
		Type,
	})
	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	return txn, nil
}

const getCreditedTransaction = `-- name: GetCreditedTransaction :many
SELECT id, client_id, amount, type, created_at FROM client_transaction
WHERE client_id=$1 AND type=$2 AND "created_at" BETWEEN NOW() - INTERVAL '1 MONTH' AND NOW()
`

type GetCreditedTransactionParams struct {
	ClientID string      `json:"client_id"`
	Type     db.Transfer `json:"type"`
}

func (b balanceStorage) GetCreditedTransaction(ctx context.Context, arg GetCreditedTransactionParams) ([]dto.ClientTransaction, error) {
	rows, err := b.db.Query(ctx, getCreditedTransaction, arg.ClientID, arg.Type)
	if err != nil {
		return nil, err
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
