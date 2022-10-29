package dto

import (
	"github.com/shopspring/decimal"
	"sms-gateway/internal/constant/model"
	"sms-gateway/internal/constant/model/db"
	"time"
)

type ClientInvoice struct {
	Id                 string               `json:"id"`
	InvoiceNumber      string               `json:"invoice_number"`
	ClientId           string               `json:"client_id"`
	CurrentBalance     decimal.Decimal      `json:"current_balance"`
	PaymentType        db.PaymentType       `json:"payment"`
	BalanceAtBeginning decimal.Decimal      `json:"balance_at_beginning"`
	MessageCount       []model.MessageCount `json:"message_count"`
	ClientTransaction  []ClientTransaction  `json:"client_transaction"`
	Discount           decimal.Decimal      `json:"discount"`
	Tax                decimal.Decimal      `json:"tax"`
	TaxRate            decimal.Decimal      `json:"tax_rate"`
	CreatedAt          *time.Time           `json:"created_at"`
}

//
//type ClientInvoice struct {
//	Id                      string          `json:"id"`
//	InvoiceNumber           string          `json:"invoice_number"`
//	PaymentType             db.PaymentType  `json:"payment_type"`
//	ClientEmail             string          `json:"client_email"`
//	BalanceAtMonthBeginning decimal.Decimal `json:"balance_at_month_beginning"`
//	CurrentBalance          decimal.Decimal `json:"current_balance"`
//	//Discount                []decimal.Decimal   `json:"discount"`
//	MessageCount       []MessageCount          `json:"message_count"`
//	ClientTransactions []dto.ClientTransaction `json:"client_transaction"`
//	Tax                decimal.Decimal         `json:"tax"`
//	TaxRate            decimal.Decimal         `json:"tax_rate"`
//	CreatedAt          *time.Time              `json:"created_at"`
//}
