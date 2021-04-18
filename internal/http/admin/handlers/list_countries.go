package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"context"
	"log"
)

const logListCountries = "[ListCountries]"

func (h *Handler) ListCountries(ctx context.Context) (*dtos.ListCountriesResponse, error) {

	result, err := h.db.ListCountries(ctx)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListCountries, err)
		return nil, err
	}

	return &dtos.ListCountriesResponse{
		Countries: result,
	}, err
}
