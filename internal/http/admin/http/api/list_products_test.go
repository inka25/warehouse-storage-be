package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListProducts(t *testing.T) {

	tts := []struct {
		caseName    string
		handlerFunc func(ctx context.Context, p *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error)
		request     func() *http.Request
		result      func(resp *http.Response)
	}{
		{
			caseName: "when all param is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/products", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error) {
				return &dtos.ListProductsResponse{
					Products: stores.Products{},
					HasNext:  false,
					Page:     1,
				}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)
			},
		},
		{
			caseName: "when there is invalid param",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/products?product_type_id=asd", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error) {
				return nil, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

				respByte, _ := json.Marshal(responseBody.Data)
				expectedByte, _ := json.Marshal([]string{errs.ErrInvalidProductTypeID.Error()})
				assert.JSONEq(t, string(expectedByte), string(respByte))
			},
		},
		{
			caseName: "when no result found",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/products?product_type_id=123", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.ListProductsRequest) (*dtos.ListProductsResponse, error) {
				return nil, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.CommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		router := mux.NewRouter()
		router.Handle("/list/products", ListProducts(tt.handlerFunc))

		rr := httptest.NewRecorder()
		req := tt.request()
		router.ServeHTTP(rr, req)

		tt.result(rr.Result())
	}

}
