// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"time"
)

type MessageType string

const (
	MessageTypeIncoming MessageType = "Incoming"
	MessageTypeOutGoing MessageType = "OutGoing"
)

func (e *MessageType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MessageType(s)
	case string:
		*e = MessageType(s)
	default:
		return fmt.Errorf("unsupported scan type for MessageType: %T", src)
	}
	return nil
}

type NullMessageType struct {
	MessageType MessageType
	Valid       bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMessageType) Scan(value interface{}) error {
	if value == nil {
		ns.MessageType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MessageType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMessageType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.MessageType, nil
}

type PaymentType string

const (
	PaymentTypePrepaid  PaymentType = "Prepaid"
	PaymentTypePostpaid PaymentType = "Postpaid"
)

func (e *PaymentType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentType(s)
	case string:
		*e = PaymentType(s)
	default:
		return fmt.Errorf("unsupported scan type for PaymentType: %T", src)
	}
	return nil
}

type NullPaymentType struct {
	PaymentType PaymentType
	Valid       bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPaymentType) Scan(value interface{}) error {
	if value == nil {
		ns.PaymentType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PaymentType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPaymentType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.PaymentType, nil
}

type Transfer string

const (
	TransferCREDITING Transfer = "CREDITING"
	TransferDEBITING  Transfer = "DEBITING"
)

func (e *Transfer) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Transfer(s)
	case string:
		*e = Transfer(s)
	default:
		return fmt.Errorf("unsupported scan type for Transfer: %T", src)
	}
	return nil
}

type NullTransfer struct {
	Transfer Transfer
	Valid    bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransfer) Scan(value interface{}) error {
	if value == nil {
		ns.Transfer, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Transfer.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransfer) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Transfer, nil
}

type Balance struct {
	ID        uuid.UUID       `json:"id"`
	ClientID  string          `json:"client_id"`
	Amount    decimal.Decimal `json:"amount"`
	Status    string          `json:"status"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
}

type Client struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Phone     []string   `json:"phone"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ClientTransaction struct {
	ID        uuid.UUID       `json:"id"`
	ClientID  string          `json:"client_id"`
	Amount    decimal.Decimal `json:"amount"`
	Type      Transfer        `json:"type"`
	CreatedAt *time.Time      `json:"created_at"`
}

type Invoice struct {
	ID                 uuid.UUID       `json:"id"`
	InvoiceNumber      uuid.NullUUID   `json:"invoice_number"`
	ClientID           string          `json:"client_id"`
	PaymentType        PaymentType     `json:"payment_type"`
	CurrentBalance     decimal.Decimal `json:"current_balance"`
	BalanceAtBeginning decimal.Decimal `json:"balance_at_beginning"`
	Discount           decimal.Decimal `json:"discount"`
	MessageCount       pgtype.JSONB    `json:"message_count"`
	ClientTransaction  pgtype.JSONB    `json:"client_transaction"`
	Tax                decimal.Decimal `json:"tax"`
	TaxRate            decimal.Decimal `json:"tax_rate"`
	CreatedAt          *time.Time      `json:"created_at"`
}

type Message struct {
	ID             uuid.UUID       `json:"id"`
	ClientID       string          `json:"client_id"`
	SenderPhone    string          `json:"sender_phone"`
	Content        string          `json:"content"`
	Price          decimal.Decimal `json:"price"`
	ReceiverPhone  string          `json:"receiver_phone"`
	Type           MessageType     `json:"type"`
	Status         string          `json:"status"`
	DeliveryStatus string          `json:"delivery_status"`
	MessageID      string          `json:"message_id"`
	CreatedAt      *time.Time      `json:"created_at"`
}

type SystemConfig struct {
	ID           uuid.UUID  `json:"id"`
	SettingName  string     `json:"setting_name"`
	SettingValue string     `json:"setting_value"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type Template struct {
	ID         uuid.UUID  `json:"id"`
	TemplateID string     `json:"template_id"`
	Client     string     `json:"client"`
	Template   string     `json:"template"`
	Category   string     `json:"category"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID  `json:"id"`
	FullName  string     `json:"full_name"`
	Phone     string     `json:"phone"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
