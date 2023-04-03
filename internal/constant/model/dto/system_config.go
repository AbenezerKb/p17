package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

type SystemConfig struct {
	ID           string     `json:"id"`
	SettingName  string     `json:"setting_name"`
	SettingValue string     `json:"setting_value"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

func (s SystemConfig) Validate() error {

	return validation.ValidateStruct(&s,
		validation.Field(&s.SettingName, validation.Required.Error("Setting Name is required")),
		validation.Field(&s.SettingValue, validation.Required.Error("Setting Value is required")),
	)
}
