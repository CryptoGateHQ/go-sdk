# cryptogate-go-sdk

Official Go SDK for the [CryptoGate](https://cryptogate.live) payment API.

## Installation

```bash
go get github.com/cryptogate/go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/cryptogate/go-sdk"
)

func main() {
    cg := cryptogate.New("sk_live_your_api_key")

    tx, err := cg.Transactions.Create(context.Background(), cryptogate.CreateParams{
        Crypto: "BTC",
        Amount: 50.00,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(tx.TXID)       // MTX-A1B2C3D4
    fmt.Println(tx.PaymentURL) // Redirect your customer here
}
```

## API Reference

```go
ctx := context.Background()

// Transactions
tx, err     := cg.Transactions.Create(ctx, cryptogate.CreateParams{Crypto: "BTC", Amount: 25.00})
tx, err     := cg.Transactions.Get(ctx, "MTX-A1B2C3D4")
result, err := cg.Transactions.List(ctx, 1, 20)
tx, err     := cg.Transactions.Cancel(ctx, "MTX-A1B2C3D4")

// Cryptos & Prices
cryptos, err := cg.Cryptos.List(ctx)
rates, err   := cg.Prices.Get(ctx)
```

## Webhook Verification

```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    sig      := r.Header.Get("X-CryptoGate-Signature")

    if !cryptogate.VerifyWebhook(os.Getenv("WEBHOOK_SECRET"), body, sig) {
        http.Error(w, "Invalid signature", http.StatusBadRequest)
        return
    }

    var event map[string]any
    json.Unmarshal(body, &event)
    // handle event["event"]
    w.WriteHeader(http.StatusOK)
}
```

## Error Handling

```go
tx, err := cg.Transactions.Get(ctx, "MTX-INVALID")
if err != nil {
    if cryptogate.IsNotFound(err) {
        fmt.Println("Transaction not found")
    } else if apiErr, ok := err.(*cryptogate.APIError); ok {
        fmt.Println(apiErr.Code, apiErr.Message)
    }
}
```

## License

MIT
