package sumaregi

import (
	"context"
	"net/http"

	"github.com/Tsubasa-2005/sumaregi-go/types"
	"github.com/google/go-querystring/query"
)

const (
	APIPathTransactions = "transactions"
)

func (c *Client) GetTransactions(ctx context.Context, opts types.GetTransactionsOpts) (*types.GetTransactionsResponse, error) {
	var result types.GetTransactionsResponse

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, APIPathTransactions, http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
