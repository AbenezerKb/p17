package invoice

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"sms-gateway/internal/adapter/storage/persistance/invoice"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/model/dto"
	"sms-gateway/internal/constant/rest/error_types"
)

type invoiceModule struct {
	invoiceStorage invoice.Storage
	common         const_init.Utils
}

//TODO add invoice method check the id for valid UUID

//Module interface for invoice module
type Module interface {
	GenerateInvoice(ctx context.Context) error
	GetClientInvoices(ctx context.Context, clientID string) ([]dto.ClientInvoice, error)
}

func InitModule(common const_init.Utils, invoiceStorage invoice.Storage) Module {
	return invoiceModule{
		invoiceStorage: invoiceStorage,
		common:         common,
	}
}

//GenerateInvoice generates invoices for all clients in monthly interval
func (im invoiceModule) GenerateInvoice(ctx context.Context) error {

	var clientInvoices []dto.ClientInvoice

	dbx := pgxpool.Conn{}
	err := dbx.BeginFunc(ctx, func(tx pgx.Tx) error {

		clients, err := im.invoiceStorage.ListClients(ctx)

		if err != nil {
			return err
		}

		for _, client := range clients {

			lastMonthBlc, err := im.invoiceStorage.GetLastMonthClientBalance(ctx, client.Id)

			if err != nil {
				return err
			}

			currentBlc, err := im.invoiceStorage.GetCurrentClientBalance(ctx, client.Id)

			if err != nil {
				return err
			}

			msgCount, err := im.invoiceStorage.LastMonthMessagesPriceAndCount(ctx, client.Id)

			if err != nil {
				return err
			}

			txn, err := im.invoiceStorage.GetLastMonthClientTransactions(ctx, client.Id)

			if err != nil {
				return err
			}

			clientInvoice := dto.ClientInvoice{
				PaymentType:             client.PaymentType,
				ClientId:                client.Id,
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
	if err != nil {
		return error_types.GetDbError(err)
	}

	return nil
}

//GetClientInvoices get the client invoice by searching using the client ID
func (im invoiceModule) GetClientInvoices(ctx context.Context, clientId string) ([]dto.ClientInvoice, error) {
	clientsInvoices, err := im.invoiceStorage.ListAllClientInvoices(ctx, clientId)
	if err != nil {
		return nil, err
	}
	return clientsInvoices, nil
}
