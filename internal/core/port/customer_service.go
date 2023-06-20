package port

import "github.com/brightnc/go-hexagonal/internal/handler/dto"

type CustomerService interface {
	GetAllUsers() ([]dto.CustomerResponse, error)
	GetUserById(id int) (*dto.CustomerResponse, error)
}
