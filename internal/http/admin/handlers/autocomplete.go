package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"context"
	"log"
)

const logAutocomplete = "[Autocomplete]"

func (h *Handler) Autocomplete(ctx context.Context, p *dtos.AutocompleteRequest) (*dtos.AutocompleteResponse, error) {

	result, err := h.db.Autocomplete(ctx, p.Prefix)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logAutocomplete, err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, errs.ErrNoResultFound
	}

	return &dtos.AutocompleteResponse{Result: result}, err
}
