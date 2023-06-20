package service_test

import (
	"testing"
	"time"

	"github.com/brightnc/go-hexagonal/internal/core/domain"
	"github.com/brightnc/go-hexagonal/internal/core/service"
	"github.com/brightnc/go-hexagonal/internal/handler/dto"
	"github.com/brightnc/go-hexagonal/mocks"
	"github.com/brightnc/go-hexagonal/pkg/errs"
	"github.com/brightnc/go-hexagonal/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsers(t *testing.T) {
	now := time.Now()
	currentTime := now.Format("2006-01-02 15:04:05")
	mockResponse := []dto.CustomerResponse{
		{ID: 1, UserName: "testuser1", Status: 1},
		{ID: 2, UserName: "testuser2", Status: 0},
		{ID: 3, UserName: "testuser3", Status: 1},
		{ID: 4, UserName: "testuser4", Status: 1},
		{ID: 5, UserName: "testuser5", Status: 0},
		{ID: 6, UserName: "testuser6", Status: 1},
	}

	mockData := []domain.Customer{
		{ID: 1, Username: "testuser1", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 2, Username: "testuser2", Password: "testPassword", Email: "test@example.com", Status: 0, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 3, Username: "testuser3", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 4, Username: "testuser4", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 5, Username: "testuser5", Password: "testPassword", Email: "test@example.com", Status: 0, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 6, Username: "testuser6", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
	}
	t.Run("success", func(t *testing.T) {
		customerMock := new(mocks.CustomerRepository)

		customerMock.On("GetAll").Return(mockData, nil).Once()

		customerService := service.NewCustomerService(customerMock)

		result, err := customerService.GetAllUsers()

		assert.NoError(t, err)
		assert.Equal(t, mockResponse, result)

		customerMock.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		logger.Init()
		customerMock := new(mocks.CustomerRepository)
		customerMock.On("GetAll").Return(nil, errs.NewUnexpectedError()).Once()

		customerService := service.NewCustomerService(customerMock)

		_, err := customerService.GetAllUsers()

		assert.Error(t, err)

		customerMock.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	now := time.Now()
	currentTime := now.Format("2006-01-02 15:04:05")

	mockResponse := []dto.CustomerResponse{
		{ID: 1, UserName: "testuser1", Status: 1},
		{ID: 2, UserName: "testuser2", Status: 0},
		{ID: 3, UserName: "testuser3", Status: 1},
		{ID: 4, UserName: "testuser4", Status: 1},
		{ID: 5, UserName: "testuser5", Status: 0},
		{ID: 6, UserName: "testuser6", Status: 1},
	}

	mockData := []domain.Customer{
		{ID: 1, Username: "testuser1", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 2, Username: "testuser2", Password: "testPassword", Email: "test@example.com", Status: 0, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 3, Username: "testuser3", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 4, Username: "testuser4", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 5, Username: "testuser5", Password: "testPassword", Email: "test@example.com", Status: 0, CreatedAt: currentTime, UpdatedAt: currentTime},
		{ID: 6, Username: "testuser6", Password: "testPassword", Email: "test@example.com", Status: 1, CreatedAt: currentTime, UpdatedAt: currentTime},
	}

	for _, d := range mockData {
		t.Run("success", func(t *testing.T) {
			customerMock := new(mocks.CustomerRepository)
			customerMock.On("GetById", mock.AnythingOfType("int")).Return(&d, nil).Once()

			customerService := service.NewCustomerService(customerMock)

			result, err := customerService.GetUserById(d.ID)

			assert.NoError(t, err)
			assert.Equal(t, &mockResponse[d.ID-1], result)

			customerMock.AssertExpectations(t)

		})
	}

	t.Run("error", func(t *testing.T) {
		logger.Init()
		customerMock := new(mocks.CustomerRepository)
		customerMock.On("GetById", mock.AnythingOfType("int")).Return(nil, errs.NewUnexpectedError()).Once()

		customerService := service.NewCustomerService(customerMock)

		_, err := customerService.GetUserById(8)

		assert.Error(t, err)

		customerMock.AssertExpectations(t)
	})

}
