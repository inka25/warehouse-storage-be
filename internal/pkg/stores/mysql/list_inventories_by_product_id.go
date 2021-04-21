package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listInventoriesByProductIdQuery = `
	SELECT w.name as warehouse, i.stock as stock
	FROM inventories i
	JOIN warehouses w ON w.id = i.warehouse_id
	WHERE product_id = ?;
`
)

func (c *Client) ListInventoriesByProductId(ctx context.Context, p *stores.SearchParams) (stores.Inventories, error) {
	var dest []stores.Inventory
	stmt, err := c.preparedStmt(listInventoriesByProductIdQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest, p.ProductID); err != nil {
		return nil, err
	}
	return dest, nil
}
