package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"net/http"
	"net/url"
)

const (
	keyValuePrefix = "prefix"
)

func Autocomplete(handlerfunc func(ctx context.Context, prefix string) (*dtos.AutocompleteResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		results, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		data, err := handlerfunc(r.Context(), results.Query().Get(keyValuePrefix))
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
