package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"log"
)

const logGetProductDetail = "[GetProductDetail]"

func (h *Handler) GetProductDetail(ctx context.Context, p *dtos.GetProductDetailRequest) (*dtos.GetProductDetailResponse, error) {

	searchParams := stores.SearchParams{
		ProductID: p.ProductId,
	}
	product, err := h.db.GetProductByProductId(ctx, &searchParams)
	if err != nil {
		log.Printf(
			"%s fail to get product err: %v\n",
			logGetProductDetail, err)
		return nil, err
	}

	histories, err := h.db.ListHistoriesByProductId(ctx, &searchParams)
	if err != nil {
		log.Printf(
			"%s fail to get histories err: %v\n",
			logGetProductDetail, err)
		return nil, err
	}

	inventories, err := h.db.ListInventoriesByProductId(ctx, &searchParams)
	if err != nil {
		log.Printf(
			"%s fail to get inventories err: %v\n",
			logGetProductDetail, err)
		return nil, err
	}

	return &dtos.GetProductDetailResponse{
		Product:     *product,
		Histories:   histories,
		Inventories: inventories,
	}, nil
}
