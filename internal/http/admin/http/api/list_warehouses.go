package api

import (
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"net/http"
)

func ListWarehouses(handlerfunc func(ctx context.Context) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := handlerfunc(r.Context())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		if data == nil {
			responder.ResponseError(w, errs.ErrNoResultFound)
			return
		}

		responder.ResponseOK(w, responder.AdvanceCommonResponse{
			Status:      0,
			Description: "success",
			Data:        data,
		})
	}
}
