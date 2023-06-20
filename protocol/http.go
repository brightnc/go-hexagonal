package protocol

import (
	"fmt"

	"github.com/brightnc/go-hexagonal/internal/handler/httphdl"
	"github.com/brightnc/go-hexagonal/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func ServeREST() {
	e := echo.New()
	custService := app.cusSvc
	custHttp := httphdl.NewCustomerHandler(custService)

	accService := app.accSvc
	accHttp := httphdl.NewAccountHandler(accService)

	v1 := e.Group("/api/v1")
	v1.GET("/customers", custHttp.GetCustomers)
	v1.GET("/customers/:id", custHttp.GetCustomerById)
	v1.POST("/customers/:id/accounts", accHttp.CreateAccount)
	v1.GET("/customers/:id/accounts", accHttp.GetAccount)

	logger.Info("Server started on port :: " + viper.GetString("app.port"))
	e.Start(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
