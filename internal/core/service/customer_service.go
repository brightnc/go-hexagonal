package service

import (
	"database/sql"

	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/brightnc/go-hexagonal/internal/handler/dto"
	"github.com/brightnc/go-hexagonal/pkg/errs"
	"github.com/brightnc/go-hexagonal/pkg/logger"
)

type customerService struct {
	repo port.CustomerRepository
}

func NewCustomerService(repo port.CustomerRepository) port.CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (r *customerService) GetAllUsers() ([]dto.CustomerResponse, error) {

	var responses []dto.CustomerResponse

	users, err := r.repo.GetAll()
	if err != nil {

		logger.Error(err)

		return nil, err
	}
	for _, user := range users {
		response := dto.CustomerResponse{
			ID:       user.ID,
			UserName: user.Username,
			Status:   user.Status,
		}
		responses = append(responses, response)
	}
	return responses, nil
}

func (r *customerService) GetUserById(id int) (*dto.CustomerResponse, error) {
	user, err := r.repo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}

		logger.Error(err)

		return nil, errs.NewUnexpectedError()
	}

	response := dto.CustomerResponse{
		ID:       user.ID,
		UserName: user.Username,
		Status:   user.Status,
	}
	return &response, nil

}
