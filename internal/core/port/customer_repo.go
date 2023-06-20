package port

import (
	"github.com/brightnc/go-hexagonal/internal/core/domain"
)

type CustomerRepository interface {
	GetAll() ([]domain.Customer, error)
	GetById(id int) (*domain.Customer, error)
	CreateUser(user domain.Customer) (*domain.Customer, error)
}
