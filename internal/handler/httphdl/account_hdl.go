package httphdl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/brightnc/go-hexagonal/internal/handler/dto"
	"github.com/brightnc/go-hexagonal/pkg/errs"
	"github.com/brightnc/go-hexagonal/pkg/logger"
	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	s port.AccountService
}

func NewAccountHandler(s port.AccountService) *accountHandler {
	return &accountHandler{
		s: s,
	}
}

func (svc *accountHandler) CreateAccount(c echo.Context) error {
	id := c.Param("id")
	idI, _ := strconv.Atoi(id)

	if c.Request().Header.Get("Content-Type") != "application/json" {
		code, msg := handleError(errs.NewValidationError("request body incorrect format"))
		return echo.NewHTTPError(code, msg)
	}

	req := dto.NewAccountRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		logger.Error(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	newAcc, err := svc.s.NewAccount(idI, req)
	if err != nil {
		code, msg := handleError(err)
		return echo.NewHTTPError(code, msg)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusCreated, newAcc)
}

func (svc *accountHandler) GetAccount(c echo.Context) error {
	id := c.Param("id")
	idI, _ := strconv.Atoi(id)

	fmt.Printf("PATH ID : %v\n", idI)

	var responses []dto.AccountResponse

	accounts, err := svc.s.GetAccounts(idI)

	if err != nil {
		code, msg := handleError(err)
		return echo.NewHTTPError(code, msg)
	}

	for _, account := range accounts {
		response := dto.AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		}

		responses = append(responses, response)
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, responses)
}
