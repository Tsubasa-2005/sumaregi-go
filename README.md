# Go API client for sumaregi

sumaregi API のクライアントライブラリ

## Usage

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/Tsubasa-2005/sumaregi-go"
	"github.com/Tsubasa-2005/sumaregi-go/domain"
	"github.com/Tsubasa-2005/sumaregi-go/types"
	"github.com/Tsubasa-2005/sumaregi-go/utils"
)

func main() {
	development := false
	envVari, err := sumaregi.LoadEnv(development)
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	config := sumaregi.NewConfig(envVari)
	scopes := []string{sumaregi.TransactionsRead}
	client, err := sumaregi.NewClient(config, scopes, envVari)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	now := time.Now()
	transactionDateTimeFrom := now.AddDate(0, 0, -2)
	transactionDateTimeTo := now.AddDate(0, 0, -1)

	transactions, err := client.GetTransactions(ctx, types.GetTransactionsOpts{
		TransactionDateTimeFrom: domain.FormatToISO8601(transactionDateTimeFrom),
		TransactionDateTimeTo:   domain.FormatToISO8601(transactionDateTimeTo),
		Limit:                   2,
	})
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	utils.PrintResponse(*transactions)
	for _, transaction := range *transactions {
		transactionDetail, err := client.GetTransactionDetail(ctx, types.GetTransactionDetailOpts{}, transaction.TransactionHeadID)
		if err != nil {
			log.Printf("Failed to get transaction details for ID %s: %v", transaction.TransactionHeadID, err)
			break
		}
		utils.PrintResponse(*transactionDetail)
	}
}

```

## References

- [sumaregi API リファレンス](https://www1.smaregi.dev/apidoc/)
- [zenn : sumaregi API クラアントライブラリ開発](https://zenn.dev/ttsbs/articles/c588ab5c1fd71f)

## APIs

### 取引

- [x] GET /{契約ID}/pos/transactions 取引一覧取得
- [x] GET /{契約ID}/pos/transactions/{id} 取引取得
