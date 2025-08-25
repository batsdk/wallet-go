package main

import (
	"log"
	"os"
	"wallet-api-go-bc/logging"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("WARNING: No .env file found or error loading it: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger,_ := zap.NewProduction()
	defer logger.Sync()

	logging.Logger = logger;

	e := echo.New()
	e.Use(middleware.Recover())


	// e.POST("/wallets", handlers.CreateWallet)
	// e.GET("/wallets/:id", handlers.GetWallet)
	// e.POST("/wallets/:id/transactions", handlers.AddTransaction)
	// e.GET("/wallets/:id/transactions", handlers.ListTransactions)

	e.Logger.Fatal(e.Start(":"+port))
}
