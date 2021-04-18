package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listbrandsQuery = `
	SELECT id, name from brands where deleted = 0;
`
)

func (c *Client) ListBrands(ctx context.Context) (stores.Results, error) {
	var dest []stores.Result
	stmt, err := c.preparedStmt(listbrandsQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}
