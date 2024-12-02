package sumaregi

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// GetProducts retrieves a list of products using the GET /products endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getItem
func (c *Client) GetProducts(ctx context.Context, opts GetProductsOpts) (*GetProductsResponse, error) {
	var result GetProductsResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, APIPathProducts, http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// PostProduct creates products using the POST /products endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/create
func (c *Client) PostProduct(ctx context.Context, params PostProductsParams) (*PostProductsResponse, error) {
	var result PostProductsResponse

	err := c.call(ctx, APIPathProducts, http.MethodPost, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
