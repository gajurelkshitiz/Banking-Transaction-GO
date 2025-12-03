package routes

import (
	"banking_transaction_go/controllers"
	"banking_transaction_go/database"
	"banking_transaction_go/middlewares"
	"banking_transaction_go/repositories"
	"banking_transaction_go/services"

	"github.com/labstack/echo/v4"
)

func BankAccountRoutes(e *echo.Echo) {
	bankRepo := repositories.NewBankAccountRepository(database.DB)
	txnRepo := repositories.NewTransactionRepository(database.DB)
	bankService := services.NewBankAccountService(bankRepo, txnRepo)
	bankController := &controllers.BankAccountController{Service: bankService}

	g := e.Group("/api/v1/accounts")
	g.POST("", bankController.Create, middlewares.AuthMiddleware)
	g.GET("", bankController.ListByUser, middlewares.AuthMiddleware)
	g.GET("/:id", bankController.Get)
	g.DELETE("/:id", bankController.Delete)
	g.POST("/:id/deposit", bankController.Deposit)
	g.POST("/:id/withdraw", bankController.Withdraw)
	g.POST("/transfer", bankController.Transfer)
}
