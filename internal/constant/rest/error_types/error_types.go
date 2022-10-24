package error_types

import "github.com/joomcode/errorx"

// traits
var (
	InvalidDate = errorx.RegisterTrait("invalid_date")

	FailedToAddSubscription    = errorx.RegisterTrait(ErrFailedToAddSubscription)
	FailedToCancelSubscription = errorx.RegisterTrait(ErrFailedToCancelSubscription)
)

var (
	ErrInvalidToken             = errorx.NewType(errorx.CommonErrors, "invalid_token")
	ErrValidationError          = errorx.NewType(errorx.CommonErrors, "invalid_input", errorx.NotFound())
	ErrValueCantBeEmpty         = errorx.NewType(errorx.CommonErrors, "error_value_can't_be_empty")
	ErrInvalidMusicCode         = errorx.NewType(errorx.CommonErrors, "invalid_input")
	ErrCancelSubscriptionFailed = errorx.NewType(errorx.CommonErrors, "invalid_input")
	ErrGenerateTokenError       = errorx.NewType(errorx.CommonErrors, "token generation error")
	ErrForbiddenMethod          = errorx.NewType(errorx.CommonErrors, "method not allowed")
	ErrDuplicateData            = errorx.AssertionFailed.NewSubtype("resource_duplicated", errorx.Duplicate())
	ErrDataNotFound             = errorx.AssertionFailed.NewSubtype("resource_not_found", errorx.NotFound())
	ErrInvalidDateArgument      = errorx.AssertionFailed.NewSubtype("error_invalid_date_argument", InvalidDate)
)
