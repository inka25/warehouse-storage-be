package mysql

import (
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
)

func (c *Client) GetProductByProductId(ctx context.Context, p *stores.SearchParams) (*stores.Product, error) {

	//var product stores.Product
	p.Limit = 1
	products, err := c.ListProducts(ctx, p)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, errs.ErrNoResultFound
	}

	return &products[0], nil
}
