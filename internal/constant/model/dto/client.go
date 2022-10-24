package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type Client struct {
	Id          string     `json:"id"`
	ClientTitle string     `json:"title"`
	Phone       string     `json:"phone"`
	Email       string     `json:"email"`
	Status      string     `json:"status"`
	Password    string     `json:"password"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (c Client) Validate() error {

	return validation.ValidateStruct(&c,
		validation.Field(&c.Phone, validation.Required.Error("Phone is required")),
		validation.Field(&c.ClientTitle, validation.Required.Error("Client Title is required")),
		validation.Field(&c.Password, validation.Required.Error("Password is required")),
		validation.Field(&c.Email, validation.Required.Error("Email is required")),
	)
}
