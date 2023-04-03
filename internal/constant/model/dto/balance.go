package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shopspring/decimal"
	"time"
)

type Balance struct {
	Id          string          `json:"id"`
	ClientId    string          `json:"client_id"`
	ClientEmail string          `json:"client_email"`
	Amount      decimal.Decimal `json:"amount"`
	Status      string          `json:"status"`
	CreatedAt   *time.Time      `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
}

func (b Balance) Validate() error {

	return validation.ValidateStruct(&b,
		validation.Field(&b.Amount, validation.Required.Error("Amount is required")),
	)
}
