package handlers

import (
	"context"
	"log"
)

const logListWarehouses = "[ListWarehouses]"

func (h *Handler) ListWarehouses(ctx context.Context) (interface{}, error) {

	result, err := h.db.ListWarehouses(ctx)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListWarehouses, err)
		return nil, err
	}

	return result, err
}
