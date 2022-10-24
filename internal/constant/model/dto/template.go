package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TemplateMessage struct {
	Template string `json:"template"`
	Category string `json:"category"`
}

type Template struct {
	Id         string `json:"id"`
	Template   string `json:"template"`
	TemplateID string `json:"template_id"`
	Client     string `json:"client"`
	Category   string `json:"category"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updated_at"`
}

func (t Template) Validate() error {

	return validation.ValidateStruct(&t,
		validation.Field(&t.Client, validation.Required.Error("Client is required")),
		validation.Field(&t.Template, validation.Required.Error("Template is required")),
	)
}
