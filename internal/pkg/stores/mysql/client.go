package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"github.com/jmoiron/sqlx"
)

type Clienter interface {
	Autocomplete(ctx context.Context, prefix string) ([]string, error)
	ListWarehouses(ctx context.Context) (stores.Results, error)
	ListProductTypes(ctx context.Context) (stores.Results, error)
	ListBrands(ctx context.Context) (stores.Results, error)
	ListProducts(ctx context.Context, p *stores.ListProductsParams) (stores.Products, error)
}

type Client struct {
	DB    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func NewClient(db *sqlx.DB) Clienter {
	return &Client{
		DB:    db,
		stmts: map[string]*sqlx.Stmt{},
	}
}

func (c *Client) preparedStmt(query string) (stmt *sqlx.Stmt, err error) {

	stmt, ok := c.stmts[query]
	if !ok {
		stmt, err = c.DB.Preparex(query)
		if err != nil {
			return nil, err
		}
		c.stmts[query] = stmt
	}
	return stmt, nil
}
