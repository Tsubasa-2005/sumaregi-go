package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

const (
	APIPathTransactions = "transactions"
)

func (c *Client) GetTransactions(ctx context.Context, opts GetTransactionsOpts) (*GetTransactionsResponse, error) {
	var result GetTransactionsResponse

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

func (c *Client) GetTransactionDetail(ctx context.Context, opts GetTransactionDetailOpts, id string) (*GetTransactionDetailResponse, error) {
	var result GetTransactionDetailResponse
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	err = c.call(ctx, path.Join(APIPathTransactions, id, "details"), http.MethodGet, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
