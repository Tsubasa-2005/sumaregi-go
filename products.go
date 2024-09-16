package sumaregi

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	APIPathProducts = "products"
)

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

func (c *Client) PostProducts(ctx context.Context, params PostProductsParams) (*PostProductsResponse, error) {
	var result PostProductsResponse

	err := c.call(ctx, APIPathProducts, http.MethodPost, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
