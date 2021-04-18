package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
)

const (
	listproductsQuery = `
	SELECT * FROM (
	SELECT p.id as id, p.code as code, p.name as name, b.name as brand, pt.name as type, sum(s.stocks) as stock
	FROM products p 
	JOIN brands b ON p.brand_id = b.id
	JOIN product_types pt ON pt.id = p.product_type_id 
	JOIN stocks s ON p.id = s.product_id
	%s
	GROUP BY p.id ) AS result
	%s	
;
`
)

func (c *Client) ListProducts(ctx context.Context, params *stores.ListProductsParams) (stores.Products, error) {
	var dest []stores.Product

	strstmt, qparams := buildStmt(params)
	stmt, err := c.preparedStmt(strstmt)

	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest, qparams...); err != nil {
		return nil, err
	}

	if len(dest) == 0 {
		return nil, errs.ErrNoResultFound
	}

	return dest, nil
}

func buildStmt(p *stores.ListProductsParams) (string, []interface{}) {

	var params []interface{}

	limitOffset := "LIMIT ? OFFSET ?"
	isDownload := p.Offset == 0 && p.Limit == 0

	if p.Prefix == "" && p.ProductTypeID == 0 && p.BrandID == 0 {
		if !isDownload {
			params = append(params, p.Limit, p.Offset)
			return fmt.Sprintf(listproductsQuery, "", limitOffset), params
		}
		return fmt.Sprintf(listproductsQuery, "", ""), params
	}

	where := "WHERE "
	andFlag := false

	if p.Prefix != "" {
		prefix := fmt.Sprintf("%%%s%%", p.Prefix)
		where = fmt.Sprintf("%s (p.code LIKE ? or p.name LIKE ?)", where)
		andFlag = true
		params = append(params, prefix, prefix)
	}
	if p.ProductTypeID != 0 {
		if andFlag {
			where = fmt.Sprintf("%s AND ", where)
		}
		where = fmt.Sprintf("%s pt.id = ?", where)
		params = append(params, p.ProductTypeID)
		if !andFlag {
			andFlag = true
		}
	}
	if p.BrandID != 0 {
		if andFlag {
			where = fmt.Sprintf("%s AND ", where)
		}
		where = fmt.Sprintf("%s b.id = ?", where)
		params = append(params, p.BrandID)
	}

	if !isDownload {
		params = append(params, p.Limit, p.Offset)
		return fmt.Sprintf(listproductsQuery, where, limitOffset), params
	}
	return fmt.Sprintf(listproductsQuery, where, ""), params
}
