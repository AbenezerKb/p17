package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/shopspring/decimal"
	"sms-gateway/internal/constant/model/db"
	"sms-gateway/internal/constant/model/dto"
	errors "sms-gateway/internal/constant/rest/error_types"
	"time"
)

type SMS struct {
	To      []string `json:"to"`
	Content string   `json:"content"`
}

type OutGoingSMS struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type Send struct {
	Sms []SMS `json:"messages"`
}

type ClientLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MessageCount struct {
	Price decimal.Decimal `json:"price"`
	Count string          `json:"count"`
	Sum   string          `json:"sum"`
}

//
//type PaymentType string
//
//const (
//	PaymentTypePostpaid PaymentType = "Postpaid"
//	PaymentTypePrepaid  PaymentType = "Prepaid"
//)

type ClientInvoice struct {
	Id                      string          `json:"id"`
	InvoiceNumber           string          `json:"invoice_number"`
	PaymentType             db.PaymentType  `json:"payment_type"`
	ClientEmail             string          `json:"client_email"`
	BalanceAtMonthBeginning decimal.Decimal `json:"balance_at_month_beginning"`
	CurrentBalance          decimal.Decimal `json:"current_balance"`
	//Discount                []decimal.Decimal   `json:"discount"`
	MessageCount       []MessageCount          `json:"message_count"`
	ClientTransactions []dto.ClientTransaction `json:"client_transaction"`
	Tax                decimal.Decimal         `json:"tax"`
	TaxRate            decimal.Decimal         `json:"tax_rate"`
	CreatedAt          *time.Time              `json:"created_at"`
}

type Transfer string

const (
	TransferCREDITING Transfer = "CREDITING"
	TransferDEBITING  Transfer = "DEBITING"
)

var TaxRate decimal.Decimal

type CustomClaims struct {
	Email string
	jwt.StandardClaims
}

func (ts *CustomClaims) GenerateToken(issuer string) (*string, error) {
	claims := &CustomClaims{
		issuer,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secretKey"))
	if err != nil {

		return nil, errors.ErrGenerateTokenError.Wrap(err, errors.ErrorGenerateTokenError)
	}

	return &t, nil
}
