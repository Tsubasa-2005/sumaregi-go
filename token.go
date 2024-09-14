package sumaregi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type AccessTokenResponse struct {
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func getAccessToken(scopes []string) (string, error) {
	requestURL := fmt.Sprintf("%s/app/%s/token", os.Getenv("SMAREGI_IDP_HOST"), os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID"))
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", os.Getenv("SMAREGI_CLIENT_ID"), os.Getenv("SMAREGI_CLIENT_SECRET"))))

	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("scope", joinScopes(scopes))

	req, err := http.NewRequest("POST", requestURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var accessTokenResult AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResult)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	return accessTokenResult.AccessToken, nil
}
