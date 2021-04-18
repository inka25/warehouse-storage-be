package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listproducttypesQuery = `
	SELECT id, name from product_types where deleted = 0;
`
)

func (c *Client) ListProductTypes(ctx context.Context) (stores.ProductTypes, error) {
	var dest []stores.ProductType
	stmt, err := c.preparedStmt(listproducttypesQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}
