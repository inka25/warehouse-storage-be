package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"context"
	"google.golang.org/appengine/log"
)

const logAutocomplete = "[Autocomplete]"

func (h *Handler) Autocomplete(ctx context.Context, prefix string) (*dtos.AutocompleteResponse, error) {

	result, err := h.db.Autocomplete(ctx, prefix)
	if err != nil {
		log.Errorf(
			ctx,
			"%s err: %v",
			logAutocomplete, err)
		return nil, err
	}

	return &dtos.AutocompleteResponse{
		Result: result,
	}, err
}
