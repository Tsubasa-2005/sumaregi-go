package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Tsubasa-2005/sumaregi-go"
	"github.com/go-resty/resty/v2"
)

type Store struct {
	StoreID   string `json:"storeId"`
	StoreName string `json:"storeName"`
}

func main() {
	sumaregi.LoadEnv()

	accessToken, err := sumaregi.GetAccessToken()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	fmt.Println("Access Token:", accessToken)

	client := resty.New()

	// Get stores
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+accessToken).
		Get(os.Getenv("SMAREGI_API_HOST") + "/" + os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID") + "/pos/stores")

	if err != nil {
		log.Fatalf("Failed to get stores: %v", err)
	}

	fmt.Println("Update Response Body:", string(resp.Body()))

	var stores []Store
	err = json.Unmarshal(resp.Body(), &stores)
	if err != nil {
		log.Fatalf("Failed to unmarshal stores response: %v", err)
	}

	fmt.Println("-----")
	fmt.Printf("Store Count: %d\n", len(stores))
	if len(stores) > 0 {
		fmt.Printf("Store ID: %s\n", stores[0].StoreID)
		fmt.Printf("Store Name: %s\n", stores[0].StoreName)
	}
}
