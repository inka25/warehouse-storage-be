package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const logListProducts = "[ListProducts]"

func (h *Handler) ListProducts(ctx context.Context, p *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error) {

	t, _ := json.MarshalIndent(p, "", " ")
	fmt.Println(string(t))

	var res dtos.ListProductsResponse
	// check and set defaultValues
	validPageLimit(p)

	t, _ = json.MarshalIndent(p, "", " ")
	fmt.Println(string(t))

	result, err := h.db.ListProducts(ctx, &stores.ListProductsParams{
		BrandID:       p.BrandID,
		ProductTypeID: p.ProductTypeID,
		Prefix:        p.Prefix,
		Page:          p.Page,
		Limit:         p.Limit,
	})
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListProducts, err)
		return nil, err
	}

	fmt.Println(len(result))
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
