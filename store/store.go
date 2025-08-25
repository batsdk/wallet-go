package store

import (
	"errors"
	"sync"

	"wallet-api-go-bc/logging"
	"wallet-api-go-bc/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	Wallets = make(map[string]*models.Wallet)
	Lock    sync.Mutex
)

func CreateWallet(name string) (*models.Wallet, error) {
	if name == "" {
		err := errors.New("wallet name is required")
		logging.Logger.Error("Failed to create wallet", zap.Error(err))
		return nil, err
	}

	Lock.Lock()
	defer Lock.Unlock()

	id := uuid.New().String()
	wallet := &models.Wallet{
		ID:      id,
		Name:    name,
		Balance: 0,
		Txns:    []models.Transaction{},
	}

	Wallets[wallet.ID] = wallet
	logging.Logger.Info("Wallet created", zap.String("wallet_id", wallet.ID), zap.String("name", name))

	return wallet, nil

}

func GetWallet(id string) (*models.Wallet, error) {

	Lock.Lock()
	defer Lock.Unlock()

	wallet, exists := Wallets[id]

	if !exists {
		err := errors.New("wallet not found")
		logging.Logger.Error("Failed to get wallet", zap.String("wallet_id", id), zap.Error(err))
		return nil, err
	}

	return wallet,nil

}

func AddTransaction(id string, txn models.Transaction) error {

	if txn.Amount <= 0 {
		err := errors.New("Transaction amount must be positive")
		return err
	}

	if txn.Type != "credit" && txn.Type != "debit" {
		err := errors.New("Transaction type must be 'credit' or 'debit'")
		return err
	}

	Lock.Lock()
	defer Lock.Unlock()

	wallet, exists := Wallets[id]
	if !exists {
		err := errors.New("wallet not found")
		logging.Logger.Error("Failed to add transaction; Wallet not found", zap.String("wallet_id", id), zap.Error(err))
		return err
	}

	if txn.Type == "debit" {
		if wallet.Balance < txn.Amount {
			err := errors.New("insufficient funds for debit transaction")
			return err
		}
		wallet.Balance -= txn.Amount
	} else {
		wallet.Balance += txn.Amount
	}

	wallet.Txns = append(wallet.Txns, txn)
	logging.Logger.Info("Transaction added", zap.String("wallet_id", id))

	return nil

}

func GetTransactions(id string, offset, limit int) (*models.PaginatedTransactionsResponse, error) {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}

	Lock.Lock()
	defer Lock.Unlock()

	wallet, exists := Wallets[id]
	if !exists {
		err := errors.New("wallet not found")
		return nil, err
	}

	txns := wallet.Txns
	total := len(txns) // Total count before pagination

	var paginatedTxns []models.Transaction

	if offset < total {
		end := offset + limit
		if end > total {
			end = total
		}
		paginatedTxns = txns[offset:end]
	} else {
		paginatedTxns = []models.Transaction{}
		logging.Logger.Debug("No transactions found for pagination", zap.String("id", id), zap.Int("offset", offset), zap.Int("limit", limit))
	}

	return &models.PaginatedTransactionsResponse{
		Transactions: paginatedTxns,
		Total:        total,
		Limit:        limit,
		Offset:       offset,
	}, nil
}