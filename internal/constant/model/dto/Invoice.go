package dto

import (
	"github.com/shopspring/decimal"
	"sms-gateway/internal/constant/model/db"
	"time"
)

type Invoice struct {
	Id            string          `json:"id"`
	InvoiceNumber int64           `json:"invoice_number"`
	Client        string          `json:"client_id"`
	Amount        decimal.Decimal `json:"amount"`
	Payment       db.Payment      `json:"payment"`
	MonthlyTotal  decimal.Decimal `json:"monthly_total"`
	Discount      decimal.Decimal `json:"discount"`
	Tax           decimal.Decimal `json:"tax"`
	TaxRate       decimal.Decimal `json:"tax_rate"`
	CreatedAt     *time.Time      `json:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at"`
}
