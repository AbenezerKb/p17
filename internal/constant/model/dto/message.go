package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shopspring/decimal"
	"sms-gateway/internal/constant/model/db"
	"time"
)

type Message struct {
	Id            string          `json:"id"`
	ReceiverPhone string          `json:"receiver_phone"`
	SenderPhone   string          `json:"sender_phone"`
	Content       string          `json:"content"`
	Price         decimal.Decimal `json:"price"`
	MsgType       db.MessageType  `json:"msg_type"`
	Status        string          `json:"status"`
	CreatedAt     *time.Time      `json:"created_at"`
}

func (m Message) Validate() error {

	return validation.ValidateStruct(&m,
		validation.Field(&m.SenderPhone, validation.Required.Error("Sender Phone is required")),
		validation.Field(&m.ReceiverPhone, validation.Required.Error("Receiver is required")),
		validation.Field(&m.Content, validation.Required.Error("Content is required")),
	)
}
