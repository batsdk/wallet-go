package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())


	// e.POST("/wallets", handlers.CreateWallet)
	// e.GET("/wallets/:id", handlers.GetWallet)
	// e.POST("/wallets/:id/transactions", handlers.AddTransaction)
	// e.GET("/wallets/:id/transactions", handlers.ListTransactions)

	e.Logger.Fatal(e.Start(":8080"))
}
