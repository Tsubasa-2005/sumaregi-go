package sumaregi

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// GetStocks retrieves stock information using the GET /stock endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getStocks
func (c *Client) GetStocks(ctx context.Context, opts GetStockOpts) (*GetStockResponse, error) {
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
