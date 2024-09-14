package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

// AccessTokenResponse represents the response from the API for the access token.
type AccessTokenResponse struct {
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

// LoadEnv loads environment variables from .env file.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// GetAccessToken fetches the access token from the API.
func GetAccessToken() (*AccessTokenResponse, error) {
	clientID := os.Getenv("SMAREGI_CLIENT_ID")
	clientSecret := os.Getenv("SMAREGI_CLIENT_SECRET")
	smaregiIDPHost := os.Getenv("SMAREGI_IDP_HOST")
	smaregiSandboxContractID := os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID")

	// Construct the URL with the contract ID
	url := fmt.Sprintf("%s/app/%s/token", smaregiIDPHost, smaregiSandboxContractID)
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret)))

	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Basic "+auth).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
			"scope":      "",
		}).
		Post(url)

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var accessTokenResult AccessTokenResponse
	err = json.Unmarshal(resp.Body(), &accessTokenResult)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &accessTokenResult, nil
}
