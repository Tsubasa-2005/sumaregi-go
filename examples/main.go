package main

import (
	"context"
	"log"

	"github.com/Tsubasa-2005/sumaregi-go"
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
	scopes := []string{sumaregi.ProductsRead, sumaregi.ProductsWrite}
	client, err := sumaregi.NewClient(config, scopes, envVari)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	products, err := client.GetProducts(ctx, types.GetProductsOpts{})
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	utils.PrintResponse(*products)
}
