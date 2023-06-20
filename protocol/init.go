package protocol

import (
	"github.com/brightnc/go-hexagonal/config"
	"github.com/brightnc/go-hexagonal/database"
	"github.com/brightnc/go-hexagonal/internal/core/port"
	"github.com/brightnc/go-hexagonal/internal/core/service"
	"github.com/brightnc/go-hexagonal/internal/repository"
	"github.com/brightnc/go-hexagonal/pkg/logger"
)

var app *application

type application struct {
	cusSvc port.CustomerService
	accSvc port.AccountService
}

func init() {
	logger.Init()
	config.InitConfig()
	db := database.InitDatabase()

	/* customerRepo : */
	customerRepo := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepo)

	accountRepo := repository.NewAccountRepositoryDB(db)
	accountService := service.NewAccountService(accountRepo)

	app = &application{
		cusSvc: customerService,
		accSvc: accountService,
	}
}
