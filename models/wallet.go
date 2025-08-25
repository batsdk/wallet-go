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
