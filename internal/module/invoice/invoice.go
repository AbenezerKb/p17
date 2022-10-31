package invoice

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"sms-gateway/internal/adapter/storage/persistance/invoice"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/rest/error_types"
)

type invoiceModule struct {
	invoiceStorage invoice.Storage
	validate       *validator.Validate
	trans          ut.Translator
}

//TODO add invoice method check the id for valid UUID

//Module interface for invoice module
type Module interface {
	GenerateInvoice(ctx context.Context) error
}

//GenerateInvoice generates invoices for all clients in monthly interval
func (im invoiceModule) GenerateInvoice(ctx context.Context) error {

	var clientInvoices []model.ClientInvoice

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
	if err != nil {
		return error_types.GetDbError(err)
	}

	return nil
}
