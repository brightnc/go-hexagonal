package httphdl

import (
	"net/http"

	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/brightnc/go-hexagonal/pkg/errs"
	"github.com/labstack/echo/v4"
)

type customerHandler struct {
	s port.CustomerService
}

func NewCustomerHandler(s port.CustomerService) *customerHandler {
	return &customerHandler{
		s: s,
	}
}

func (svc *customerHandler) GetCustomers(c echo.Context) error {
	users, err := svc.s.GetAllUsers()
	if err != nil {
		return echo.ErrInternalServerError
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(200)
	return c.JSON(http.StatusOK, users)
}

func (svc *customerHandler) GetCustomerById(c echo.Context) error {

	var id int

	if err := echo.PathParamsBinder(c).Int("id", &id).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errs.NewValidationError("invalid path parameter"))
	}

	c.Response().Header().Set("Content-Type", "application/json")

	user, err := svc.s.GetUserById(id)
	if err != nil {
		code, msg := handleError(err)
		return echo.NewHTTPError(code, msg)
	}

	c.Response().WriteHeader(200)
	return c.JSON(http.StatusOK, user)
}
