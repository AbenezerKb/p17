package balance

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	const_init "sms-gateway/internal/constant/init"
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
	ListAllBalance(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error)
	GetClientBalance(ctx context.Context, phone string) (*dto.Balance, error)
	UpdateClientBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error)
}

func StorageInit(utils const_init.Utils) Storage {
	return balanceStorage{
		db:  utils.Conn,
		dbp: *db.New(utils.Conn),
	}
}

func (c balanceStorage) AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {

	blc, err := c.dbp.AddBalance(ctx, db.AddBalanceParams{
		ClientID: balance.Client,
		Amount:   balance.Amount,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)

	}

	balance = &dto.Balance{
		Id:        blc.ID.String(),
		Client:    blc.ClientID,
		Amount:    blc.Amount,
		Status:    blc.Status,
		CreatedAt: blc.CreatedAt,
	}

	return balance, nil
}

func (b balanceStorage) ListAllBalance(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error) {
	page, _ := strconv.ParseInt(params.Page, 10, 32)
	perPage, _ := strconv.ParseInt(params.PerPage, 10, 32)

	resizedPage := int32(page)
	resizedPerPage := int32(perPage)
	blc, err := b.dbp.ListAllBalance(ctx, db.ListAllBalanceParams{
		Offset: resizedPage,
		Limit:  resizedPerPage,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	var balance []dto.Balance

	for _, v := range blc {
		balance = append(balance, dto.Balance{
			Id:        v.ID.String(),
			Client:    v.ClientID,
			Amount:    v.Amount,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return balance, nil
}

func (b balanceStorage) GetClientBalance(ctx context.Context, phone string) (*dto.Balance, error) {

	blc, err := b.dbp.GetClientBalance(ctx, phone)

	if err != nil {
		return nil, error_types.GetDbError(err)
	}

	balance := &dto.Balance{
		Id:        blc.ID.String(),
		Client:    blc.ClientID,
		Amount:    blc.Amount,
		Status:    blc.Status,
		CreatedAt: blc.CreatedAt,
		UpdatedAt: blc.UpdatedAt,
	}

	return balance, nil
}

func (b balanceStorage) UpdateClientBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {

	updatedBlc, err := b.dbp.UpdateClientBalance(ctx, db.UpdateClientBalanceParams{
		ClientID: balance.Client,
		Amount:   balance.Amount,
	})

	if err != nil {
		return nil, error_types.GetDbError(err)
	}
	updatedBalance := &dto.Balance{
		Id:        updatedBlc.ID.String(),
		Client:    updatedBlc.ClientID,
		Amount:    updatedBlc.Amount,
		Status:    updatedBlc.Status,
		CreatedAt: updatedBlc.CreatedAt,
		UpdatedAt: updatedBlc.UpdatedAt,
	}

	return updatedBalance, nil

}
