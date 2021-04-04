package handlers

import (
	"context"
	"log"
)

const logListProducts = "[ListProducts]"

func (h *Handler) ListProducts(ctx context.Context) (interface{}, error) {

	result, err := h.db.ListProducts(ctx)
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListWarehouses, err)
		return nil, err
	}

	return result, err
}

