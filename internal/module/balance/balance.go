package balance

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"sms-gateway/internal/adapter/storage/persistance/balance"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest"
)

type balanceModule struct {
	balanceStorage balance.Storage
	validate       *validator.Validate
	trans          ut.Translator
}

type Module interface {
	AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error)
	ListAllBalance(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error)
	GetClientBalance(ctx context.Context, email string) (*dto.Balance, error)
	UpdateClientBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error)
}

func ModuleInit(balanceStorage balance.Storage, utils const_init.Utils) Module {
	return balanceModule{
		balanceStorage: balanceStorage,
		validate:       utils.GoValidator,
		trans:          utils.Translator,
	}
}

func (b balanceModule) AddBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {

	blc, err := b.balanceStorage.AddBalance(ctx, balance)
	if err != nil {
		return nil, err
	}

	return blc, nil
}

func (b balanceModule) ListAllBalance(ctx context.Context, params *rest.QueryParams) ([]dto.Balance, error) {
	balanceList, err := b.ListAllBalance(ctx, params)
	if err != nil {
		return nil, err
	}

	return balanceList, nil
}

func (b balanceModule) GetClientBalance(ctx context.Context, email string) (*dto.Balance, error) {

	blc, err := b.balanceStorage.GetCurrentClientBalance(ctx, email)
	if err != nil {
		return nil, err
	}
	return blc, nil
}

func (b balanceModule) UpdateClientBalance(ctx context.Context, balance *dto.Balance) (*dto.Balance, error) {
	blc, err := b.balanceStorage.UpdateClientBalance(ctx, model.TransferCREDITING, balance)
	if err != nil {
		return nil, err
	}
	return blc, nil
}

func (b balanceModule) GenerateInvoice(ctx context.Context) error {

	var clientInvoices []model.ClientInvoice

	dbx := pgxpool.Conn{}
	err := dbx.BeginFunc(ctx, func(tx pgx.Tx) error {

		clients, err := b.balanceStorage.ListClients(ctx)

		if err != nil {
			return err
		}

		for _, client := range clients {

			lastMonthBlc, err := b.balanceStorage.GetLastMonthClientBalance(ctx, client.Id)

			if err != nil {
				return err
			}

			currentBlc, err := b.balanceStorage.GetCurrentClientBalance(ctx, client.Id)

			if err != nil {
				return err
			}

			msgCount, err := b.balanceStorage.LastMonthMessagesPriceAndCount(ctx, client.Id)

			if err != nil {
				return err
			}

			txn, err := b.balanceStorage.GetLastMonthClientTransactions(ctx, client.Id, db.TransferCREDITING)

			if err != nil {
				return err
			}

			clientInvoice := model.ClientInvoice{
				PaymentType:             client.PaymentType,
				ClientEmail:             client.Email,
				BalanceAtMonthBeginning: lastMonthBlc.Amount,
				CurrentBalance:          currentBlc.Amount,
				MessageCount:            msgCount,
				ClientTransactions:      txn,
				//TODO get the tax rate from config
				Tax:     decimal.NewFromInt(0),
				TaxRate: decimal.NewFromInt(0),
			}
			clientInvoices = append(clientInvoices, clientInvoice)
		}
		return nil
	})
	b.balanceStorage.A(ctx, clientInvoices)
	//blc, err := b.balanceStorage.UpdateClientBalance(ctx, balance)
	//if err != nil {
	//	return nil, err
	//}
	//return blc, nil
	//TODO implementation detail
	return nil
}
