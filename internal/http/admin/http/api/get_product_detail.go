package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"net/http"
	"net/url"
	"strconv"
)

const (
	keyValueId = "id"
)

func GetProductDetail(handlerfunc func(ctx context.Context, p *dtos.GetProductDetailRequest) (*dtos.GetProductDetailResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestUrl, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		results := requestUrl.Query()
		param := dtos.GetProductDetailRequest{}
		param.ProductId, err = strconv.ParseInt(results.Get(keyValueId), 10, 64)
		if err != nil {
			responder.ResponseError(w, errs.ErrInvalidId)
			return
		}

		data, err := handlerfunc(r.Context(), &param)
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
