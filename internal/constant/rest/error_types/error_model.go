package error_types

type ErrorModel struct {
	// ErrorCode        string      `json:"errorCode"`
	ErrorDescription string      `json:"errorDescription"`
	ErrorMessage     string      `json:"errorMessage"`
	ValidationErrors interface{} `json:"fieldErrors,omitempty"`
}
