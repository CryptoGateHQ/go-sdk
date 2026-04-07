package cryptogate

import (
	"context"
	"net/http"
)

// CryptosService handles supported cryptocurrencies.
type CryptosService struct{ c *Client }

// Crypto represents a supported cryptocurrency.
type Crypto struct {
	Symbol     string `json:"symbol"`
	Name       string `json:"name"`
	Blockchain string `json:"blockchain"`
	Network    string `json:"network"`
	Type       string `json:"type"`
}

// CryptosResult holds the list of supported cryptos.
type CryptosResult struct {
	Cryptocurrencies []Crypto `json:"cryptocurrencies"`
	Total            int      `json:"total"`
}

// List returns all cryptocurrencies supported for payments.
func (s *CryptosService) List(ctx context.Context) (*CryptosResult, error) {
	var result CryptosResult
	err := s.c.do(ctx, http.MethodGet, "/cryptos/list", nil, &result)
	return &result, err
}

// PricesService handles exchange rates.
type PricesService struct{ c *Client }

// RatesResult holds exchange rate data.
type RatesResult struct {
	Rates     map[string]float64 `json:"rates"`
	Currency  string             `json:"currency"`
	FetchedAt string             `json:"fetched_at"`
}

// Get returns current USD exchange rates for all supported cryptos.
func (s *PricesService) Get(ctx context.Context) (*RatesResult, error) {
	var result RatesResult
	err := s.c.do(ctx, http.MethodGet, "/prices", nil, &result)
	return &result, err
}
