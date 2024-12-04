package sumaregi_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Tsubasa-2005/sumaregi-go"
	"github.com/stretchr/testify/require"
)

func TestClient_GetTransactionCSV(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	envVari, err := sumaregi.LoadEnv(false)
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	config := sumaregi.NewConfig(envVari)
	scopes := []string{sumaregi.TransactionsRead}
	client, err := sumaregi.NewClient(config, scopes, envVari)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	now := time.Now()
	transactionDateTimeFrom := now.AddDate(0, 0, -2)
	transactionDateTimeTo := now.AddDate(0, 0, -1)

	t.Run("success", func(t *testing.T) {
		csv, err := client.PostTransactionCSV(ctx, sumaregi.PostTransactionCSVOpts{
			TransactionDateTimeFrom: sumaregi.FormatToISO8601(transactionDateTimeFrom),
			TransactionDateTimeTo:   sumaregi.FormatToISO8601(transactionDateTimeTo),
			CallbackURL:             "https://example.com/callback",
		})
		require.NoError(t, err)

		fmt.Print(csv)
	})
}
