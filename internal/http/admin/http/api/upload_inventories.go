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

func UploadInventories(handlerfunc func(ctx context.Context, p *dtos.UploadInventoriesRequest) (error, []string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestUrl, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		results := requestUrl.Query()
		param := dtos.UploadInventoriesRequest{}
		var errors []string
		param.WarehouseID, err = strconv.ParseInt(results.Get(keyValueWarehouseId), 10, 64)
		if err != nil {
			errors = append(errors, errs.ErrInvalidWarehouseID.Error())
		}

		if len(errors) > 0 {
			responder.ResponseError(w, errs.ErrInvalidRequestParam, errors)
			return
		}

		err, errs := handlerfunc(r.Context(), &param)
		if err != nil {
			responder.ResponseError(w, err, errs)
			return
		}

		responder.ResponseOK(w, responder.CommonResponse{
			Status:      0,
			Description: "success",
		})
	}
}

func UploadInventoriesTemplate(handlerfunc func(ctx context.Context) (*dtos.UploadInventoriesTemplateResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := handlerfunc(r.Context())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseCSVDownload(w, res.Filename, res.UploadInventories)
	}
}
