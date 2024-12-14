package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

// GetCategory retrieves specific category
// using the GET /categories/{id} endpoint.
// https://www1.smaregi.dev/apidoc/#operation/categoryGetId
func (c *Client) GetCategory(ctx context.Context, opts GetCategoryOpts, id string) (*GetCategoryResponse, error) {
	var result GetCategoryResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, path.Join(APIPathCategories, id), http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
