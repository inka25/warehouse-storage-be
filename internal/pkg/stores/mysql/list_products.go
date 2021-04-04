package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

const (
	listproductsQuery = `
	SELECT p.id as id, p.code as code, p.name as name, b.name as brand, pt.name as type, sum(s.stocks) as stock
FROM products p 
JOIN brands b ON p.brand_id = b.id
JOIN product_types pt ON pt.id = p.product_type_id 
JOIN stocks s ON p.id = s.product_id 
WHERE 
GROUP BY p.id;

`
)

func (c *Client) ListProducts(ctx context.Context) (stores.Products, error) {
	var dest []stores.Product
	stmt, err := c.preparedStmt(listproductsQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}


func getStmt(searchBy )