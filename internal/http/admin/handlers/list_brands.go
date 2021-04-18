package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"context"
	"log"
)

const logListBrands = "[ListBrands]"

func (h *Handler) ListBrands(ctx context.Context) (*dtos.ListBrandsResponse, error) {

	result, err := h.db.ListBrands(ctx)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListBrands, err)
		return nil, err
	}

	return &dtos.ListBrandsResponse{
		Brands: result,
	}, err
}
