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
	keyValueProductTypeId = "product_type_id"
	keyValueBrandId = "brand_id"
	keyValuePage = "page"
	keyValueLimit = "limit"
)


func ListProducts(handlerfunc func(ctx context.Context, param *dtos.ListProductsRequest) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {


		results, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}
		var param *dtos.ListProductsRequest
		var errors []error
		param.Auto = results.Query().Get(keyValuePrefix)
		param.WarehouseID, err = strconv.ParseInt(results.Query().Get(keyValueWarehouseId), 10, 64)
if err != nil{
	errors = append(errors, errs.ErrInvalidWarehouseID)
}
		param.ProductTypeID, err = strconv.ParseInt(results.Query().Get(keyValueProductTypeId), 10, 64)
		if err != nil{
			errors = append(errors, errs.ErrInvalidProductTypeID)
		}
		param.BrandID, err = strconv.ParseInt(results.Query().Get(keyValueBrandId), 10, 64)
		if err != nil{
			errors = append(errors, errs.ErrInvalidBrandID)
		}
		param.Page, err = strconv.ParseInt(results.Query().Get(keyValuePage), 10, 64)
		if err != nil{
			errors = append(errors, errs.ErrInvalidPageNumber)
		}
		param.Limit, err = strconv.ParseInt(results.Query().Get(keyValueLimit), 10, 64)
		if err != nil{
			errors = append(errors, errs.ErrInvalidPageLimit)
		}

		if len(errors) > 0{
			responder.ResponseError(w, errs.ErrInvalidRequestParam, errors)
			return
		}

		data, err := handlerfunc(r.Context(), param)
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
