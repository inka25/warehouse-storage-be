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
	keyValueCountryId     = "country_id"
	keyValueProductTypeId = "product_type_id"
	keyValueBrandId       = "brand_id"
	keyValuePage          = "page"
	keyValueLimit         = "limit"
)

func ListProducts(handlerfunc func(ctx context.Context, param *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestUrl, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		results := requestUrl.Query()
		param := dtos.ListProductsRequest{}
		var errors []string
		param.Prefix = results.Get(keyValuePrefix)
		param.CountryID, err = strconv.ParseInt(results.Get(keyValueCountryId), 10, 64)
		if err != nil && results.Get(keyValueCountryId) != "" {
			errors = append(errors, errs.ErrInvalidProductTypeID.Error())
		}
		param.ProductTypeID, err = strconv.ParseInt(results.Get(keyValueProductTypeId), 10, 64)
		if err != nil && results.Get(keyValueProductTypeId) != "" {
			errors = append(errors, errs.ErrInvalidProductTypeID.Error())
		}
		param.BrandID, err = strconv.ParseInt(results.Get(keyValueBrandId), 10, 64)
		if err != nil && results.Get(keyValueBrandId) != "" {
			errors = append(errors, errs.ErrInvalidBrandID.Error())
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
