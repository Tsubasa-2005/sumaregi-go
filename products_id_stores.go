package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

// GetProductsIDStores retrieves store-specific details for a given product ID
// using the GET /products/{id}/stores endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getProductStores
func (c *Client) GetProductsIDStores(ctx context.Context, opts GetProductIDStoresOpts, id string) (*GetProductIDStoresResponse, error) {
	var result GetProductIDStoresResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathProducts, id, "stores"), http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
