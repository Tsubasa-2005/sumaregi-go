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
	config := sumaregi.NewConfig(envVari.SmaregiAPIHost, envVari.SmaregiContractID)
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
	})
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	utils.PrintResponse(*transactions)

}
