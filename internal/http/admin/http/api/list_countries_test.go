package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCountries(t *testing.T) {

	tts := []struct {
		caseName    string
		handlerFunc func(ctx context.Context) (*dtos.ListCountriesResponse, error)
		request     func() *http.Request
		result      func(resp *http.Response)
	}{
		{
			caseName: "when return result is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/countries", nil)
				return req
			},
			handlerFunc: func(ctx context.Context) (*dtos.ListCountriesResponse, error) {
				return &dtos.ListCountriesResponse{
					Countries: stores.Results{},
				}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)
			},
		},
		{
			caseName: "when no result found, is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/countries", nil)
				return req
			},
			handlerFunc: func(ctx context.Context) (*dtos.ListCountriesResponse, error) {
				return nil, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)
			},
		},
		{
			caseName: "when error occurs",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/list/countries", nil)
				return req
			},
			handlerFunc: func(ctx context.Context) (*dtos.ListCountriesResponse, error) {
				return nil, errors.New("")
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		router := mux.NewRouter()
		router.Handle("/list/countries", ListCountries(tt.handlerFunc))

		rr := httptest.NewRecorder()
		req := tt.request()
		router.ServeHTTP(rr, req)

		tt.result(rr.Result())
	}

}
