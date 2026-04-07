package cryptogate

import (
	"context"
	"fmt"
	"net/http"
)

// TransactionsService handles transaction API calls.
type TransactionsService struct{ c *Client }

// CreateParams are the parameters for creating a transaction.
type CreateParams struct {
	Crypto   string  `json:"crypto"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// Transaction represents a CryptoGate payment transaction.
type Transaction struct {
	TXID            string  `json:"txid"`
	Status          string  `json:"status"`
	AmountUSD       float64 `json:"amount_usd"`
	AmountPaid      float64 `json:"amount_paid"`
	AmountRemaining float64 `json:"amount_remaining"`
	Currency        string  `json:"currency"`
	Payment         Payment `json:"payment"`
	PaymentURL      string  `json:"payment_url"`
	ExpiresAt       string  `json:"expires_at"`
	CreatedAt       string  `json:"created_at"`
}

// Payment holds payment-specific details within a Transaction.
type Payment struct {
	Crypto                string  `json:"crypto"`
	AmountUSD             float64 `json:"amount_usd"`
	AmountCrypto          string  `json:"amount_crypto"`
	LockedRate            float64 `json:"locked_rate"`
	Address               string  `json:"address"`
	Confirmations         int     `json:"confirmations"`
	ConfirmationsRequired int     `json:"confirmations_required"`
	Status                string  `json:"status"`
}

// ListResult holds a paginated list of transactions.
type ListResult struct {
	Transactions []Transaction `json:"transactions"`
	Pagination   Pagination    `json:"pagination"`
}

// Pagination holds paging metadata.
type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// Create creates a new payment transaction.
func (s *TransactionsService) Create(ctx context.Context, params CreateParams) (*Transaction, error) {
	if params.Crypto == "" {
		return nil, &APIError{Code: "VALIDATION_ERROR", Message: "crypto is required", Status: 400}
	}
	if params.Amount == 0 {
		return nil, &APIError{Code: "VALIDATION_ERROR", Message: "amount is required", Status: 400}
	}
	params.Currency = "USD"
	var tx Transaction
	err := s.c.do(ctx, http.MethodPost, "/transactions/create", params, &tx)
	return &tx, err
}

// Get retrieves a transaction by ID.
func (s *TransactionsService) Get(ctx context.Context, txid string) (*Transaction, error) {
	if txid == "" {
		return nil, &APIError{Code: "VALIDATION_ERROR", Message: "txid is required", Status: 400}
	}
	var tx Transaction
	err := s.c.do(ctx, http.MethodGet, "/transactions/"+txid, nil, &tx)
	return &tx, err
}

// List returns paginated transactions for the account.
func (s *TransactionsService) List(ctx context.Context, page, limit int) (*ListResult, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	var result ListResult
	err := s.c.do(ctx, http.MethodGet, fmt.Sprintf("/transactions/list?page=%d&limit=%d", page, limit), nil, &result)
	return &result, err
}

// Cancel cancels a pending transaction.
func (s *TransactionsService) Cancel(ctx context.Context, txid string) (*Transaction, error) {
	if txid == "" {
		return nil, &APIError{Code: "VALIDATION_ERROR", Message: "txid is required", Status: 400}
	}
	var tx Transaction
	err := s.c.do(ctx, http.MethodPost, "/transactions/"+txid+"/cancel", nil, &tx)
	return &tx, err
}
