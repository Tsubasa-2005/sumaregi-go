package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

func (c *Client) GetProductsID(ctx context.Context, opts GetProductsIDOpts, id string) (*GetProductsIDResponse, error) {
	var result GetProductsIDResponse

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
