package handlers

import (
	"net/http"
	"strconv"
	"wallet-api-go-bc/logging"
	"wallet-api-go-bc/models"
	"wallet-api-go-bc/store"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func CreateWallet(c echo.Context) error {

	var req models.CreateWalletRequest

	if err := c.Bind(&req); err != nil {
		logging.Logger.Error("Failed to bind request", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	wallet, err := store.CreateWallet(req.Name)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create wallet"})
	}

	return c.JSON(http.StatusCreated, wallet)

}

func GetWallet(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wallet ID is required"})
	}

	wallet, err := store.GetWallet(id)

	if err != nil {
		logging.Logger.Warn("Get wallet failed in handler", zap.String("id", id), zap.Error(err))
		if err.Error() == "wallet not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal error"})
	}

	return c.JSON(http.StatusOK, wallet)

}

func AddTransaction(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wallet ID is required"})
	}

	var req models.CreateTransactionRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	txn := models.Transaction{
		Type:  req.Type,
		Amount: req.Amount,
	}

	err := store.AddTransaction(id, txn)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedWallet,_ := store.GetWallet(id)
	return c.JSON(http.StatusCreated, updatedWallet)

}

func ListTransactions(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wallet ID is required"})
	}

	offsetStr := c.QueryParam("offset")
	limitStr := c.QueryParam("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	resp, err := store.GetTransactions(id,offset,limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)

}