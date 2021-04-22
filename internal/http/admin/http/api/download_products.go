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

func DownloadProducts(handlerfunc func(ctx context.Context, param *dtos.DownloadProductRequest) (*dtos.DownloadProductsResponse, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestUrl, err := url.ParseRequestURI(r.URL.String())
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		results := requestUrl.Query()
		param := dtos.DownloadProductRequest{}
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

		if len(errors) > 0 {
			responder.ResponseError(w, errs.ErrInvalidRequestParam, errors)
			return
		}

		res, err := handlerfunc(r.Context(), &param)
		if err != nil {
			responder.ResponseError(w, err)
			return
		}

		responder.ResponseCSVDownload(w, res.Filename, res.Products)
	}
}
