package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listwarehousesQuery = `
	SELECT id, name from warehouses;
`
)

func (c *Client) ListWarehouses(ctx context.Context) (stores.Warehouses, error) {
	var dest []stores.Warehouse
	stmt, err := c.preparedStmt(listwarehousesQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}