package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type User struct {
	Id        string     `json:"id"`
	Phone     string     `json:"phone"`
	FullName  string     `json:"full_name"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u User) Validate() error {

	return validation.ValidateStruct(&u,
		validation.Field(&u.FullName, validation.Required.Error("Name is required")),
		validation.Field(&u.Phone, validation.Required.Error("Phone is required")),
		validation.Field(&u.Password, validation.Required.Error("Password is required"), validation.Length(8, 0)),
	)
}
