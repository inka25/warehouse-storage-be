package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
	"log"
	"time"
)

const logDownloadProducts = "[DownloadProducts]"

func (h *Handler) DownloadProducts(ctx context.Context, p *dtos.DownloadProductRequest) (*dtos.DownloadProductsResponse, error) {
	var res dtos.DownloadProductsResponse

	// for download, no need offset-limit
	result, err := h.db.ListProducts(ctx, &stores.ListProductsParams{
		BrandID:       p.BrandID,
		ProductTypeID: p.ProductTypeID,
		Prefix:        p.Prefix,
	})
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logDownloadProducts, err)
		return nil, err
	}

	res.Filename = filenameBuilder(p, &result[0])
	res.Products = result

	return &res, err
}

func filenameBuilder(p *dtos.DownloadProductRequest, res *stores.Product) string {

	date := time.Now().Format("Monday_02_January_2006")
	suffix := fmt.Sprintf("ProductsList_%s", date)
	if p.Prefix != "" {
		suffix = suffix + fmt.Sprintf("_prefix_%s", p.Prefix)
	}
	if p.BrandID != 0 {
		suffix = suffix + fmt.Sprintf("_brand_%s", res.Brand)
	}
	if p.ProductTypeID != 0 {
		suffix = suffix + fmt.Sprintf("_product_type_%s", res.Type)
	}

	return suffix

}
