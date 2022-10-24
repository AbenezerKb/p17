package error_types

const (
	ErrorInvalidDataType            = "error invalid data type"
	ErrorInvalidQueries             = "invalid queries"
	ErrorInvalidRequestBody         = "invalid request body"
	ErrMusicCodeNotFound            = "invalid music code"
	ErrFailedToCancelSubscription   = "failed to cancel subscription"
	ErrFailedToAddSubscription      = "failed to add subscription"
	ErrFailedToAddSubscriptionTwice = "You already have active subscription"
	ErrUserSubscriptionPending      = "user subscription pending"
	ErrorGenerateTokenError         = "generate token error"
	ErrorUnauthorizedError          = "unauthorized error"
)
