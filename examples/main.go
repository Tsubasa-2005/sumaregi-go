package main

import (
	"context"
	"log"
	"time"

	"github.com/Tsubasa-2005/sumaregi-go"
)

func main() {
	development := false
	envVari, err := LoadEnv(development)
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
	transactionDateTimeFrom := now.AddDate(0, 0, -30)
	transactionDateTimeTo := now.AddDate(0, 0, -1)

	transactions, err := client.GetTransactions(ctx, sumaregi.GetTransactionsOpts{
		TransactionDateTimeFrom: sumaregi.FormatToISO8601(transactionDateTimeFrom),
		TransactionDateTimeTo:   sumaregi.FormatToISO8601(transactionDateTimeTo),
	})
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	sumaregi.PrintResponse(*transactions)
	for _, transaction := range *transactions {
		transactionDetail, err := client.GetTransactionDetail(ctx, sumaregi.GetTransactionDetailOpts{}, transaction.TransactionHeadID)
		if err != nil {
			log.Printf("Failed to get transaction details for ID %s: %v", transaction.TransactionHeadID, err)
			break
		}
		sumaregi.PrintResponse(*transactionDetail)
	}
}
