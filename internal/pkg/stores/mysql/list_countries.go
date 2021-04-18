package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listcountriesQueries = `
	SELECT id, country as name from countries where deleted = 0;
`
)

func (c *Client) ListCountries(ctx context.Context) (stores.Results, error) {
	var dest []stores.Result
	stmt, err := c.preparedStmt(listcountriesQueries)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}
