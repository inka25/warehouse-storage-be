package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"context"
	"log"
)

const logListWarehouses = "[ListWarehouses]"

func (h *Handler) ListWarehouses(ctx context.Context) (*dtos.ListWareshousesResponse, error) {

	result, err := h.db.ListWarehouses(ctx)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListWarehouses, err)
		return nil, err
	}

	return &dtos.ListWareshousesResponse{
		Warehouses: result,
	}, err
}
