// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
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

type Payment string

const (
	PaymentPrepaid  Payment = "Prepaid"
	PaymentPostpaid Payment = "Postpaid"
)

func (e *Payment) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Payment(s)
	case string:
		*e = Payment(s)
	default:
		return fmt.Errorf("unsupported scan type for Payment: %T", src)
	}
	return nil
}

type NullPayment struct {
	Payment Payment
	Valid   bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPayment) Scan(value interface{}) error {
	if value == nil {
		ns.Payment, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Payment.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPayment) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Payment, nil
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
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Invoice struct {
	ID            uuid.UUID       `json:"id"`
	InvoiceNumber int64           `json:"invoice_number"`
	ClientID      string          `json:"client_id"`
	Payment       Payment         `json:"payment"`
	Amount        decimal.Decimal `json:"amount"`
	TotalMonthly  decimal.Decimal `json:"total_monthly"`
	Discount      decimal.Decimal `json:"discount"`
	Tax           decimal.Decimal `json:"tax"`
	TaxRate       decimal.Decimal `json:"tax_rate"`
	CreatedAt     *time.Time      `json:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at"`
}

type Message struct {
	ID             uuid.UUID       `json:"id"`
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
