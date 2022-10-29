package middleware

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joomcode/errorx"
	"net/http"
	const_init "sms-gateway/internal/constant/init"
	"sms-gateway/internal/constant/rest"
	"sms-gateway/internal/constant/rest/error_types"
	"strings"
)

type errorHandlingMiddleWare struct {
	utils const_init.Utils
}

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			err := err.Unwrap()
			// tempErr := errorx.Cast(err)
			status := http.StatusInternalServerError
			if errorx.HasTrait(err, errorx.Timeout()) {
				status = http.StatusRequestTimeout

			}
			if errorx.HasTrait(err, errorx.Duplicate()) {
				status = http.StatusBadRequest
			}
			if errorx.HasTrait(err, error_types.InvalidDate) {
				status = http.StatusBadRequest
			}

			// if errorx.HasTrait(err.Unwrap(), error_types.ErrCancelSubscriptionFailed) {
			// 	status = http.StatusBadRequest
			// }
			if errorx.HasTrait(err, errorx.NotFound()) {

				status = http.StatusBadRequest
			}

			// if errorx.IsOfType(err.Unwrap(), error_types.ErrInvalidMusicCode) {

			// 	status = http.StatusBadRequest
			// }

			if errorx.IsOfType(err, errorx.IllegalArgument) {

				status = http.StatusBadRequest
			}

			if errorx.IsOfType(err, error_types.ErrCancelSubscriptionFailed) {

				status = http.StatusBadRequest
			}

			if errorx.IsOfType(err, error_types.ErrGenerateTokenError) {

				status = http.StatusInternalServerError
			}

			if errorx.IsOfType(err, error_types.ErrInvalidToken) {
				status = http.StatusUnauthorized
			}
			if e, ok := err.(validation.Errors); ok {
				errResponse := &error_types.ErrorModel{
					ErrorMessage:     err.Error()[strings.LastIndex(err.Error(), ":")+2 : len(err.Error())-1],
					ErrorDescription: err.Error(),
					ValidationErrors: e,
				}
				c.JSON(http.StatusBadRequest, rest.ErrData{
					Error: errResponse,
				})
				return
			}

			if e, ok := err.(*errorx.Error); ok {
				//	fmt.Printf("%+v", e)
				c.JSON(status, rest.ErrData{

					Error: error_types.ErrorModel{
						ErrorMessage:     e.Message(),
						ErrorDescription: e.Error(),
					},
				})
			} else {
				c.JSON(status, rest.ErrData{
					Error: error_types.ErrorModel{
						ErrorMessage:     err.Error(),
						ErrorDescription: err.Error(),
					},
				})
			}
		}
	}
}
