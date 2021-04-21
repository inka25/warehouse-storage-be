package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
)

const (
	listproductsQuery = `
	SELECT * FROM (
	SELECT p.id as id, p.code as code, c.country as country, p.name as name, b.name as brand, pt.name as type, sum(i.stock) as stock
	FROM products p 
	JOIN brands b ON p.brand_id = b.id
	JOIN product_types pt ON pt.id = p.product_type_id 
	JOIN inventories i ON p.id = i.product_id
    JOIN countries c ON p.country_id = c.id
	%s
	GROUP BY p.id ) AS result
	%s	
;
`
)

func (c *Client) ListProducts(ctx context.Context, params *stores.SearchParams) (stores.Products, error) {
	var dest []stores.Product

	strstmt, qparams := buildProductQueryStmt(params)
	stmt, err := c.preparedStmt(strstmt)

	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest, qparams...); err != nil {
		return nil, err
	}

	return dest, nil
}

func buildProductQueryStmt(p *stores.SearchParams) (string, []interface{}) {

	var params []interface{}

	where := "WHERE p.deleted = 0"
	limitOffset := "LIMIT ? OFFSET ?"
	isDownload := p.Offset == 0 && p.Limit == 0

	if p.Prefix == "" && p.ProductTypeID == 0 && p.BrandID == 0 {
		if !isDownload {
			params = append(params, p.Limit, p.Offset)
			return fmt.Sprintf(listproductsQuery, where, limitOffset), params
		}
		return fmt.Sprintf(listproductsQuery, where, ""), params
	}

	if p.Prefix != "" {
		prefix := fmt.Sprintf("%%%s%%", p.Prefix)
		where = fmt.Sprintf("%s AND (p.code LIKE ? OR p.name LIKE ?)", where)
		params = append(params, prefix, prefix)
	}
	if p.ProductTypeID != 0 {
		where = fmt.Sprintf("%s AND pt.id = ?", where)
		params = append(params, p.ProductTypeID)
	}
	if p.BrandID != 0 {
		where = fmt.Sprintf("%s AND b.id = ?", where)
		params = append(params, p.BrandID)
	}
	if p.CountryID != 0 {
		where = fmt.Sprintf("%s AND c.id = ?", where)
		params = append(params, p.CountryID)
	}

	if !isDownload {
		params = append(params, p.Limit, p.Offset)
		return fmt.Sprintf(listproductsQuery, where, limitOffset), params
	}
	return fmt.Sprintf(listproductsQuery, where, ""), params
}
