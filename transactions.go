package sumaregi

import (
	"context"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

// GetTransactions retrieves transaction using the GET /transactions endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getTransaction
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

// GetTransactionDetail retrieves detailed information for a specific transaction
// using the GET /transactions/{id}/details endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/getTransactionDetail
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

// PostTransactionCSV request creation of a CSV file containing transaction details
// using the POST /transactions/details/out_file_async endpoint.
// See Smaregi API documentation for details:
// https://www1.smaregi.dev/apidoc/#operation/postCreateTransactionDetailList
func (c *Client) PostTransactionCSV(ctx context.Context, params PostTransactionCSVOpts) (*PostTransactionCSVResponse, error) {
	var result PostTransactionCSVResponse

	err := c.call(ctx, path.Join(APIPathTransactions, "details", "out_file_async"), http.MethodPost, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
