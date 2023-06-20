package port

import "github.com/brightnc/go-hexagonal/internal/handler/dto"

type AccountService interface {
	NewAccount(customerId int, acc dto.NewAccountRequest) (*dto.AccountResponse, error)
	GetAccounts(customerId int) ([]dto.AccountResponse, error)
}
