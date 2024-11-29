package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

// GetProduct retrieves specific product
// using the GET /products/{id} endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getItemId
func (c *Client) GetProduct(ctx context.Context, opts GetProductOpt, id string) (*GetProductIDStoresResponse, error) {
	var result GetProductIDStoresResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathProducts, id), http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
