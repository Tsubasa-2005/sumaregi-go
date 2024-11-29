package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

// GetStoresID retrieves details for a specific store using the GET /stores/{id} endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getStoresId
func (c *Client) GetStoresID(ctx context.Context, opts GetStoresIDOpts, id string) (*GetStoresIDResponse, error) {
	var result GetStoresIDResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathStores, id), http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
