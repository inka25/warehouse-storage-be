package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"github.com/jmoiron/sqlx"
)

type Clienter interface {
	Autocomplete(ctx context.Context, prefix string) ([]string, error)

	GetProductByProductId(ctx context.Context, params *stores.SearchParams) (*stores.Product, error)

	ListWarehouses(ctx context.Context) (stores.Results, error)
	ListProductTypes(ctx context.Context) (stores.Results, error)
	ListBrands(ctx context.Context) (stores.Results, error)
	ListCountries(ctx context.Context) (stores.Results, error)
	ListProducts(ctx context.Context, p *stores.SearchParams) (stores.Products, error)

	ListInventoriesByWarehouseId(ctx context.Context, p *stores.SearchParams) (stores.Inventories, error)
	ListInventoriesByProductId(ctx context.Context, p *stores.SearchParams) (stores.Inventories, error)

	ListHistoriesByProductId(ctx context.Context, p *stores.SearchParams) (stores.Histories, error)

	InsertInventoriesByWarehouseId(ctx context.Context, p *stores.UploadParams) error
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
