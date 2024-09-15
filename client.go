package sumaregi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	httpClient *http.Client
	config     *Config
	token      string
}

const (
	APIPathPos = "/pos"
)

func NewClient(config *Config, scopes []string, envVari EnvironmentVariable) (*Client, error) {
	token, err := getAccessToken(scopes, envVari)
	if err != nil {
		return nil, err
	}
	client := &Client{
		httpClient: &http.Client{},
		config:     config,
		token:      token,
	}
	return client, nil
}

func (c *Client) call(
	ctx context.Context,
	apiPath string, method string,
	queryParams url.Values, postBody interface{},
	res interface{},
) error {
	var (
		contentType string
		body        io.Reader
	)
	if method != http.MethodDelete {
		contentType = "application/json"
		jsonParams, err := json.Marshal(postBody)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(jsonParams)
	}

	req, err := c.newRequest(ctx, apiPath, method, contentType, queryParams, body)
	if err != nil {
		return err
	}
	return c.do(ctx, req, res)
}

func (c *Client) newRequest(
	ctx context.Context,
	apiPath string, method string,
	contentType string,
	queryParams url.Values,
	body io.Reader,
) (*http.Request, error) {
	// construct url
	u, err := url.Parse(c.config.APIEndpoint)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, c.config.ContractID, APIPathPos, apiPath)
	u.RawQuery = queryParams.Encode()
	// request with context
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	fmt.Print(u.String())
	// set http headers
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	return req, nil
}

func (c *Client) do(
	ctx context.Context,
	req *http.Request,
	res interface{},
) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if res == nil {
		return nil
	}

	if response.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
		fmt.Printf("HTTP %d: %s", response.StatusCode, string(bodyBytes))

		return fmt.Errorf("HTTP %d: %s", response.StatusCode, string(bodyBytes))
	}
	return json.NewDecoder(response.Body).Decode(&res)
}
