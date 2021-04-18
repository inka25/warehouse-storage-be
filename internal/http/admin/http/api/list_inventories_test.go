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

func TestListInventories(t *testing.T) {

	tts := []struct {
		caseName    string
		handlerFunc func(ctx context.Context, p *dtos.ListInventoriesRequest) (*dtos.ListInventoriesResponse, error)
		request     func() *http.Request
		result      func(resp *http.Response)
	}{
		{
			caseName: "when all param is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/inventories?warehouse_id=1", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.ListInventoriesRequest) (*dtos.ListInventoriesResponse, error) {
				return &dtos.ListInventoriesResponse{
					Inventories: stores.Inventories{
						{
							ID: 1,
						},
					},
					HasNext: false,
					Page:    1,
				}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)
			},
		},
		{
			caseName: "when warehouse id is empty",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/inventories", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.ListInventoriesRequest) (*dtos.ListInventoriesResponse, error) {
				return nil, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusBadRequest)

				respByte, _ := json.Marshal(responseBody.Data)
				expectedByte, _ := json.Marshal([]string{errs.ErrInvalidWarehouseID.Error()})
				assert.JSONEq(t, string(expectedByte), string(respByte))
			},
		},
		{
			caseName: "when no result found",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/inventories?warehouse_id=123", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.ListInventoriesRequest) (*dtos.ListInventoriesResponse, error) {
				return nil, errs.ErrNoResultFound
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
		router.Handle("/list/inventories", ListInventories(tt.handlerFunc))

		rr := httptest.NewRecorder()
		req := tt.request()
		router.ServeHTTP(rr, req)

		tt.result(rr.Result())
	}

}
