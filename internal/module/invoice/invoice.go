package invoice

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"sms-gateway/internal/adapter/storage/persistance/invoice"
)

type invoiceModule struct {
	invoiceStorage invoice.Storage
	validate       *validator.Validate
	trans          ut.Translator
}
