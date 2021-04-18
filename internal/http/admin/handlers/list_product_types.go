package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"context"
	"log"
)

const logListProductTypes = "[ListProductTypes]"

func (h *Handler) ListProductTypes(ctx context.Context) (*dtos.ListProductTypesResponse, error) {

	result, err := h.db.ListProductTypes(ctx)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListProductTypes, err)
		return nil, err
	}

	return &dtos.ListProductTypesResponse{
		ProductTypes: result,
	}, err
}
