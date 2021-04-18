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

func ListInventories(handlerfunc func(ctx context.Context, param *dtos.ListInventoriesRequest) (*dtos.ListInventoriesResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestUrl, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		results := requestUrl.Query()
		param := dtos.ListInventoriesRequest{}
		var errors []string
		param.WarehouseID, err = strconv.ParseInt(results.Get(keyValueWarehouseId), 10, 64)
		if err != nil {
			errors = append(errors, errs.ErrInvalidWarehouseID.Error())
		}
		param.Page, err = strconv.ParseInt(results.Get(keyValuePage), 10, 64)
		if err != nil && results.Get(keyValuePage) != "" {
			errors = append(errors, errs.ErrInvalidPageNumber.Error())
		}
		param.Limit, err = strconv.ParseInt(results.Get(keyValueLimit), 10, 64)
		if err != nil && results.Get(keyValueLimit) != "" {
			errors = append(errors, errs.ErrInvalidPageLimit.Error())
		}

		if len(errors) > 0 {
			responder.ResponseError(w, errs.ErrInvalidRequestParam, errors)
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
