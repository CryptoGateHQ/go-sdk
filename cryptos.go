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

// RatesResult holds exchange rate data returned by GET /prices.
type RatesResult struct {
	Crypto    map[string]float64 `json:"crypto"`     // USD price per coin: {"BTC": 50000, ...}
	Fiat      map[string]float64 `json:"fiat"`       // USD per fiat unit: {"USD": 1.0, "PLN": 0.25, "EUR": 1.09, "GBP": 1.27}
	FetchedAt string             `json:"fetched_at"`
}

// Get returns current USD exchange rates for all supported cryptos.
func (s *PricesService) Get(ctx context.Context) (*RatesResult, error) {
	var result RatesResult
	err := s.c.do(ctx, http.MethodGet, "/prices", nil, &result)
	return &result, err
}
