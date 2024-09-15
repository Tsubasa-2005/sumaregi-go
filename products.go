package sumaregi

import (
	"context"
	"net/http"

	"github.com/Tsubasa-2005/sumaregi-go/types"
	"github.com/google/go-querystring/query"
)

const (
	APIPathProducts = "products"
)

func (c *Client) GetProducts(ctx context.Context, opts types.GetProductsOpts) (*types.GetProductsResponse, error) {
	var result types.GetProductsResponse

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

func (c *Client) PostProducts(ctx context.Context, params types.PostProductsParams) (*types.PostProductsResponse, error) {
	var result types.PostProductsResponse

	err := c.call(ctx, APIPathProducts, http.MethodPost, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
