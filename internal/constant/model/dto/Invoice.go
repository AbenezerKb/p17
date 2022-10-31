package dto

import (
	"github.com/shopspring/decimal"
	"sms-gateway/internal/constant/model"
	"time"
)

type PaymentType string

const (
	PaymentTypePostpaid PaymentType = "Postpaid"
	PaymentTypePrepaid  PaymentType = "Prepaid"
)

type MessageCount struct {
	Price decimal.Decimal `json:"price"`
	Count string          `json:"count"`
	Sum   string          `json:"sum"`
}

type ClientTransaction struct {
	Id        string          `json:"id"`
	ClientId  string          `json:"client_id"`
	Amount    decimal.Decimal `json:"amount"`
	Type      model.Transfer  `json:"type"`
	CreatedAt *time.Time      `json:"created_at"`
}

type ClientInvoice struct {
	Id                      string          `json:"id"`
	InvoiceNumber           string          `json:"invoice_number"`
	PaymentType             PaymentType     `json:"payment_type"`
	ClientEmail             string          `json:"client_email"`
	BalanceAtMonthBeginning decimal.Decimal `json:"balance_at_month_beginning"`
	CurrentBalance          decimal.Decimal `json:"current_balance"`
	//Discount                []decimal.Decimal   `json:"discount"`
	MessageCount       []MessageCount      `json:"message_count"`
	ClientTransactions []ClientTransaction `json:"client_transaction"`
	Tax                decimal.Decimal     `json:"tax"`
	TaxRate            decimal.Decimal     `json:"tax_rate"`
	CreatedAt          *time.Time          `json:"created_at"`
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
