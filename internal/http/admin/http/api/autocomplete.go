package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"net/http"
	"net/url"
)

const (
	keyValuePrefix = "prefix"
)

func Autocomplete(handlerfunc func(ctx context.Context, p *dtos.AutocompleteRequest) (*dtos.AutocompleteResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var p dtos.AutocompleteRequest
		results, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		p.Prefix = results.Query().Get(keyValuePrefix)

		resp, err := handlerfunc(r.Context(), &p)
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseOK(w, responder.AdvanceCommonResponse{
			Status:      0,
			Description: "success",
			Data:        resp,
		})
	}
}
