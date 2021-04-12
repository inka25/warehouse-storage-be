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

func TestListWarehouses(t *testing.T) {

	tts := []struct {
		caseName    string
		handlerFunc func(ctx context.Context) (*dtos.ListWareshousesResponse, error)
		request     func() *http.Request
		result      func(resp *http.Response)
	}{
		{
			caseName: "when return result is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/warehouses", nil)
				return req
			},
			handlerFunc: func(ctx context.Context) (*dtos.ListWareshousesResponse, error) {
				return &dtos.ListWareshousesResponse{
					Warehouses: stores.Warehouses{},
				}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)
			},
		},
		{
			caseName: "when no result found",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/warehouses", nil)
				return req
			},
			handlerFunc: func(ctx context.Context) (*dtos.ListWareshousesResponse, error) {
				return nil, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
			},
		},
		{
			caseName: "when error occurs",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/warehouses", nil)
				return req
			},
			handlerFunc: func(ctx context.Context) (*dtos.ListWareshousesResponse, error) {
				return nil, errs.ErrNoResultFound
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		router := mux.NewRouter()
		router.Handle("/list/warehouses", ListWarehouses(tt.handlerFunc))

		rr := httptest.NewRecorder()
		req := tt.request()
		router.ServeHTTP(rr, req)

		tt.result(rr.Result())
	}

}
