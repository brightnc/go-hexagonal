package port

import "github.com/brightnc/go-hexagonal/internal/core/domain"

type AccountRepository interface {
	Create(domain.Account) (*domain.Account, error)
	GetAll(customerId int) ([]domain.Account, error)
}
