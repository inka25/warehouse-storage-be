package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"log"
)

const logListProducts = "[ListProducts]"

func (h *Handler) ListProducts(ctx context.Context, p *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error) {
	var res dtos.ListProductsResponse
	// check and set defaultValues
	validPageLimit(p)

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit + 1

	result, err := h.db.ListProducts(ctx, &stores.ListProductsParams{
		BrandID:       p.BrandID,
		ProductTypeID: p.ProductTypeID,
		Prefix:        p.Prefix,
		Offset:        offset,
		Limit:         limit,
	})
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListProducts, err)
		return nil, err
	}

	if len(result) > int(p.Limit) {
		res.HasNext = true
		result = result[:p.Limit]
	}
	res.Page = p.Page
	res.Products = result

	return &res, err
}

func validPageLimit(p *dtos.ListProductsRequest) {

	if p.Limit == 0 {
		p.Limit = 20
	}

	if p.Page == 0 {
		p.Page = 1
	}

}
