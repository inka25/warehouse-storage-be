package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
)

const (
	listinventoriesByWarehouseIdQuery = `
	SELECT p.id as id, w.name as warehouse, c.country as country, p.code as code, p.name as name, b.name as brand, pt.name as type, i.stock as stock
	FROM products p 
	JOIN brands b ON p.brand_id = b.id
	JOIN countries c ON p.country_id = c.id
	JOIN product_types pt ON pt.id = p.product_type_id 
	JOIN inventories i ON p.id = i.product_id
	JOIN warehouses w ON w.id = i.warehouse_id
	WHERE w.id = ? AND p.deleted = 0
	%s	
;
`
)

func (c *Client) ListInventoriesByWarehouseId(ctx context.Context, p *stores.SearchParams) (stores.Inventories, error) {
	var dest []stores.Inventory

	strstmt, qparams := buildInventoryQueryStmt(p)
	stmt, err := c.preparedStmt(strstmt)

	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest, qparams...); err != nil {
		return nil, err
	}

	return dest, nil
}

func buildInventoryQueryStmt(p *stores.SearchParams) (string, []interface{}) {

	params := []interface{}{p.WarehouseID}

	limitOffset := "LIMIT ? OFFSET ?"
	isDownload := p.Offset == 0 && p.Limit == 0

	if !isDownload {
		params = append(params, p.Limit, p.Offset)
		return fmt.Sprintf(listinventoriesByWarehouseIdQuery, limitOffset), params
	}
	return fmt.Sprintf(listinventoriesByWarehouseIdQuery, ""), params
}
