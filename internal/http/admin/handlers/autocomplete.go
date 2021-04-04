package handlers

import (
	"context"
	"log"
)

const logAutocomplete = "[Autocomplete]"

func (h *Handler) Autocomplete(ctx context.Context, prefix string) (interface{}, error) {

	result, err := h.db.Autocomplete(ctx, prefix)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logAutocomplete, err)
		return nil, err
	}

	return result, err
}
