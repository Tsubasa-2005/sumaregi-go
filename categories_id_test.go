package sumaregi_test

import (
	"context"
	"log"
	"testing"

	"github.com/Tsubasa-2005/sumaregi-go"
	"github.com/stretchr/testify/require"
)

func TestClient_GetCategory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	envVari, err := sumaregi.LoadEnv(false)
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	config := sumaregi.NewConfig(envVari)
	scopes := []string{sumaregi.ProductsRead}
	client, err := sumaregi.NewClient(config, scopes, envVari)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	product, err := client.GetProducts(ctx, sumaregi.GetProductsOpts{
		Limit: 1,
	})
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		category, err := client.GetCategory(ctx, sumaregi.GetCategoryOpts{}, (*product)[0].CategoryID)
		require.NoError(t, err)

		require.Equal(t, (*product)[0].CategoryID, category.CategoryID)
	})
}
