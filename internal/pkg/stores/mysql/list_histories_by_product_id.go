package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listHistoriesByProductIdQuery = `
	SELECT type as edit_type, description, updated_at, updated_by 
	FROM histories 
	WHERE product_id = ? LIMIT ?;
`
)

func (c *Client) ListHistoriesByProductId(ctx context.Context, p *stores.SearchParams) (stores.Histories, error) {
	var dest []stores.History
	p.Limit = 10
	stmt, err := c.preparedStmt(listHistoriesByProductIdQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest, p.ProductID, p.Limit); err != nil {
		return nil, err
	}
	return dest, nil
}
