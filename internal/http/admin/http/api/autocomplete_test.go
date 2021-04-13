package api

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/http/responder"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAutocomplete(t *testing.T) {

	tts := []struct {
		caseName    string
		handlerFunc func(ctx context.Context, p *dtos.AutocompleteRequest) (*dtos.AutocompleteResponse, error)
		request     func() *http.Request
		result      func(resp *http.Response)
	}{
		{
			caseName: "when is ok",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/auto?prefix=hello", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.AutocompleteRequest) (*dtos.AutocompleteResponse, error) {
				return &dtos.AutocompleteResponse{
					[]string{"helloyou"},
				}, nil
			},
			result: func(resp *http.Response) {
				var responseBody *responder.AdvanceCommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, resp.StatusCode, http.StatusOK)

				respByte, _ := json.Marshal(responseBody.Data)
				expectedByte, _ := json.Marshal(dtos.AutocompleteResponse{
					[]string{"helloyou"},
				})
				assert.JSONEq(t, string(expectedByte), string(respByte))

			},
		},
		{
			caseName: "when result not found",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/auto?prefix=asd", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.AutocompleteRequest) (*dtos.AutocompleteResponse, error) {
				return nil, errs.ErrNoResultFound
			},
			result: func(resp *http.Response) {
				var responseBody *responder.CommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

				assert.Error(t, errs.ErrNoResultFound)
			},
		},
		{
			caseName: "when database error",
			request: func() *http.Request {

				req, _ := http.NewRequest(http.MethodGet, "/auto?prefix=asd", nil)
				return req
			},
			handlerFunc: func(ctx context.Context, p *dtos.AutocompleteRequest) (*dtos.AutocompleteResponse, error) {
				return nil, errors.New("")
			},
			result: func(resp *http.Response) {
				var responseBody *responder.CommonResponse
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

				assert.Error(t, errs.ErrInternalServerError)
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		router := mux.NewRouter()
		router.Handle("/auto", Autocomplete(tt.handlerFunc))

		rr := httptest.NewRecorder()
		req := tt.request()
		router.ServeHTTP(rr, req)

		tt.result(rr.Result())
	}

}
