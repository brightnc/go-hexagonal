package service

import (
	"strings"
	"time"

	"github.com/brightnc/go-hexagonal/internal/core/domain"
	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/brightnc/go-hexagonal/internal/handler/dto"
	"github.com/brightnc/go-hexagonal/pkg/errs"
	"github.com/brightnc/go-hexagonal/pkg/logger"
)

type accountService struct {
	repo port.AccountRepository
}

func NewAccountService(repo port.AccountRepository) port.AccountService {
	return &accountService{
		repo: repo,
	}
}

func (r *accountService) NewAccount(customerId int, req dto.NewAccountRequest) (*dto.AccountResponse, error) {

	//Validate the request

	if req.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5000")
	}

	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type must be saving or checking")
	}

	now := time.Now()
	current_time := now.Format("2006-01-02 15:04:05")
	newAccount := domain.Account{
		CustomerID:  customerId,
		OpeningDate: current_time,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}
	acc, err := r.repo.Create(newAccount)
	if err != nil {
		logger.Error(err)

		return nil, errs.NewUnexpectedError()
	}

	accRes := dto.AccountResponse{
		AccountID:   acc.AccountID,
		OpeningDate: acc.OpeningDate,
		AccountType: acc.AccountType,
		Amount:      acc.Amount,
		Status:      acc.Status,
	}

	return &accRes, nil
}
func (r *accountService) GetAccounts(customerId int) ([]dto.AccountResponse, error) {
	var responses []dto.AccountResponse
	accounts, err := r.repo.GetAll(customerId)

	if err != nil {
		logger.Error(err)

		return nil, errs.NewNotFoundError("account not found")
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

	return responses, nil
}
