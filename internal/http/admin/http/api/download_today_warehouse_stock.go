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

func DownloadTodayStock(handlerfunc func(ctx context.Context, warehouseID int64) (*dtos.DownloadTodayStockResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		results, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		warehouseID, err := strconv.ParseInt(results.Query().Get(keyValueWarehouseId), 10, 64)
		if err != nil {
			responder.ResponseError(w, errs.ErrInvalidID)
			return
		}

		result, err := handlerfunc(r.Context(), warehouseID)
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		if result == nil {
			responder.ResponseError(w, errs.ErrNoResultFound)
			return
		}

		responder.ResponseOK(w, responder.AdvanceCommonResponse{
			Status:      0,
			Description: "success",
		})

		responder.ResponseCSVDownload(w, result.Filename, result.Data)
	}
}