package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
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

func TestGetProductDetail(t *testing.T) {

	tts := []struct {
		caseName    string
		handlerFunc func(ctx context.Context, p *dtos.GetProductDetailRequest) (*dtos.GetProductDetailResponse, error)
		request     func() *http.Request
		result      func(resp *http.Response)
	}{
		{
			caseName: "when return result is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/product?id=1", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, resp *dtos.GetProductDetailRequest) (*dtos.GetProductDetailResponse, error) {
				return &dtos.GetProductDetailResponse{
					Product: stores.Product{ID: int64(1)},
					Histories: stores.Histories{
						{
							UpdatedBy: "admin_test",
						},
					},
					Inventories: stores.Inventories{
						{
							Warehouse: "test",
						},
					},
				}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)
				assert.ObjectsAreEqualValues(dtos.GetProductDetailResponse{
					Product: stores.Product{ID: int64(1)},
					Histories: stores.Histories{
						{
							UpdatedBy: "admin_test",
						},
					},
					Inventories: stores.Inventories{
						{
							Warehouse: "test",
						},
					},
				}, responseBody.Data)
			},
		},
		{
			caseName: "when no parameter",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/product", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, resp *dtos.GetProductDetailRequest) (*dtos.GetProductDetailResponse, error) {
				return &dtos.GetProductDetailResponse{}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.CommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
			},
		},
		{
			caseName: "when bad parameter",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/product?id=asd", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, resp *dtos.GetProductDetailRequest) (*dtos.GetProductDetailResponse, error) {
				return &dtos.GetProductDetailResponse{}, nil
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
		router.Handle("/product", GetProductDetail(tt.handlerFunc))

		rr := httptest.NewRecorder()
		req := tt.request()
		router.ServeHTTP(rr, req)

		tt.result(rr.Result())
	}

}
