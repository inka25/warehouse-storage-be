package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
	"strings"
)

const (
	insertInventoriesbyWarehouseQuery = `
	INSERT INTO inventories (product_id, warehouse_id, stock, update_time, updated_by) 
	VALUES    
`
	whereOnDuplicateUpdateQuery = `
	ON DUPLICATE KEY UPDATE stock=VALUES(stock), update_time=VALUES(update_time), updated_by=VALUES(updated_by);
`
)

func (c *Client) InsertInventoriesByWarehouseId(ctx context.Context, p *stores.UploadParams) error {

	query := insertInventoriesbyWarehouseQuery
	var vals []interface{}
	for _, row := range p.UploadInfo {
		query += "(?, ?, ?, ?, ?),"
		vals = append(vals, row.ID, p.WarehouseId, row.Stock, p.Timestamp, p.Email)
	}

	query = strings.TrimSuffix(query, ",")
	query += whereOnDuplicateUpdateQuery

	fmt.Println(query)

	stmt, err := c.preparedStmt(query)
	if err != nil {
		return err
	}
	if _, err = stmt.ExecContext(ctx, vals...); err != nil {
		return err
	}
	return nil
}
