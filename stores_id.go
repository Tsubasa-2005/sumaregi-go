package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

func (c *Client) GetStoresID(ctx context.Context, opts GetStoresIDOpts, id string) (*GetStoresIDResponse, error) {
	var result GetStoresIDResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathProducts, "stores", id), http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
