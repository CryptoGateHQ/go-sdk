// Package cryptogate provides a client for the CryptoGate payment API.
//
// Usage:
//
//	cg := cryptogate.New("sk_live_your_api_key")
//	tx, err := cg.Transactions.Create(context.Background(), cryptogate.CreateParams{
//	    Crypto: "BTC",
//	    Amount: 50.00,
//	})
package cryptogate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://api.cryptogate.live"
	sdkVersion     = "1.0.0"
)

// Client is the CryptoGate API client.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client

	Transactions *TransactionsService
	Cryptos      *CryptosService
	Prices       *PricesService
}

// New creates a new CryptoGate client.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, o := range opts {
		o(c)
	}
	c.Transactions = &TransactionsService{c}
	c.Cryptos = &CryptosService{c}
	c.Prices = &PricesService{c}
	return c
}

// Option configures the client.
type Option func(*Client)

// WithBaseURL overrides the API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) { c.baseURL = url }
}

// WithTimeout sets the HTTP request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.httpClient.Timeout = d }
}

func (c *Client) do(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", "cryptogate-go-sdk/"+sdkVersion)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &APIError{Code: "NETWORK_ERROR", Message: err.Error()}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return json.Unmarshal(data, out)
	}

	var apiErr APIError
	_ = json.Unmarshal(data, &apiErr)
	apiErr.Status = resp.StatusCode
	if apiErr.Message == "" {
		apiErr.Message = fmt.Sprintf("HTTP %d", resp.StatusCode)
	}
	return &apiErr
}
