package models

type Wallet struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Balance float64       `json:"balance"`
	Txns    []Transaction `json:"transactions"`
}

type Transaction struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type CreateWalletRequest struct {
	Name string `json:"name"`
}

type CreateTransactionRequest struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type PaginatedTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
	Total        int           `json:"total"`
	Limit        int           `json:"limit"`
	Offset       int           `json:"offset"`
}