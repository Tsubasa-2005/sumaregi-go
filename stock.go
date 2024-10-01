package sumaregi

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	APIPathStock = "stock"
)

func (c *Client) GetStock(ctx context.Context, opts GetStockOpts) (*GetStockResponse, error) {
	var result GetStockResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, APIPathStock, http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
