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
	keyValueWarehouseId = "warehouse_id"
)

func DownloadInventories(handlerfunc func(ctx context.Context, p *dtos.DownloadInventoriesRequest) (*dtos.DownloadInventoriesResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestUrl, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		results := requestUrl.Query()
		param := dtos.DownloadInventoriesRequest{}
		var errors []string
		param.WarehouseID, err = strconv.ParseInt(results.Get(keyValueWarehouseId), 10, 64)
		if err != nil {
			errors = append(errors, errs.ErrInvalidWarehouseID.Error())
		}

		if len(errors) > 0 {
			responder.ResponseError(w, errs.ErrInvalidRequestParam, errors)
			return
		}

		result, err := handlerfunc(r.Context(), &param)
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseCSVDownload(w, result.Filename, result.Inventories)
	}
}
