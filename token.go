package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type AccessTokenResponse struct {
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

var (
	token          string
	tokenExpiresAt time.Time
	mu             sync.Mutex
)

// fetchNewAccessToken fetches a new access token from the API.
func fetchNewAccessToken() (*AccessTokenResponse, error) {
	clientID := os.Getenv("SMAREGI_CLIENT_ID")
	clientSecret := os.Getenv("SMAREGI_CLIENT_SECRET")
	smaregiIDPHost := os.Getenv("SMAREGI_IDP_HOST")
	smaregiSandboxContractID := os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID")

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

	fmt.Println("Response Body:", string(resp.Body()))

	var accessTokenResult AccessTokenResponse
	err = json.Unmarshal(resp.Body(), &accessTokenResult)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &accessTokenResult, nil
}

// GetAccessToken returns a valid access token, either from cache or by fetching a new one.
func GetAccessToken() (string, error) {
	mu.Lock()
	defer mu.Unlock()

	newToken, err := fetchNewAccessToken()
	if err != nil {
		return "", err
	}

	token = newToken.AccessToken

	err = SaveToEnv("ACCESS_TOKEN", token)
	if err != nil {
		return "", fmt.Errorf("error saving token to .env file: %w", err)
	}

	return token, nil
}
