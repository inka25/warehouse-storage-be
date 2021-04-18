package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"net/http"
)

func ListProductTypes(handlerfunc func(ctx context.Context) (*dtos.ListProductTypesResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := handlerfunc(r.Context())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseOK(w, responder.AdvanceCommonResponse{
			Status:      0,
			Description: "success",
			Data:        data,
		})
	}
}
