package sumaregi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type AccessTokenResponse struct {
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

// fetchNewAccessToken fetches a new access token from the API.
func getAccessToken(scopes []string) (string, error) {
	url := fmt.Sprintf("%s/app/%s/token", os.Getenv("SMAREGI_IDP_HOST"), os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID"))
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("SMAREGI_CLIENT_ID"), os.Getenv("SMAREGI_CLIENT_SECRET"))))

	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", "Basic "+auth).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
			"scope":      joinScopes(scopes),
		}).
		Post(url)

	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}

	var accessTokenResult AccessTokenResponse
	err = json.Unmarshal(resp.Body(), &accessTokenResult)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	return accessTokenResult.AccessToken, nil
}
