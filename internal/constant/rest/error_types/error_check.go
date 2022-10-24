package error_types

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func GetDbError(err error) error {

	if err.Error() == pgx.ErrNoRows.Error() {
		// return ErrDataNotFound.New(err.Error())
		return ErrDataNotFound.New("resource not found")
	}

	pqe, ok := err.(*pgconn.PgError)
	fmt.Println("the pg error ", ok)
	if !ok {
		return nil
	}
	fmt.Println("the pg error ", pqe.Error())

	if pqe.Code == "23505" {
		message := strings.ReplaceAll(strings.ReplaceAll(pqe.Detail, ")=(", " : "), "Key ", "")
		return ErrDuplicateData.New(message)
	}
	return pqe
}

func ValidationError(err error) *ErrorModel {
	if e, ok := err.(validation.Errors); ok {
		return &ErrorModel{
			ErrorMessage:     "there are errors in field validation",
			ErrorDescription: "fix validation errors and submit again",
			ValidationErrors: e,
		}
	}
	return nil
}
