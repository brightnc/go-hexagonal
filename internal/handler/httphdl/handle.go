package httphdl

import (
	"net/http"

	"github.com/brightnc/go-hexagonal/pkg/errs"
)

func handleError(err error) (code int, message string) {
	appErr, ok := err.(errs.AppError)
	if ok {
		return appErr.Code, appErr.Message
	}
	return http.StatusInternalServerError, err.Error()

}
