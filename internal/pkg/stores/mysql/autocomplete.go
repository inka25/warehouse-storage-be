package mysql

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	autocompleteQuery = `
	SELECT result
		FROM (
			SELECT code AS result 
			FROM products 
			WHERE code 
			LIKE ?) AS A 
		UNION (
			SELECT name AS result 
			FROM products 
			WHERE name 
			LIKE ?) ;
`
)

func (c *Client) Autocomplete(ctx context.Context, prefix string) ([]string, error) {
	prefix = fmt.Sprintf("%%%s%%", prefix)
	var dest []string
	stmt, err := c.preparedStmt(autocompleteQuery)
	if err != nil {
		return nil, err
	}
	if err = stmt.SelectContext(ctx, &dest, prefix, prefix); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return dest, nil
}
