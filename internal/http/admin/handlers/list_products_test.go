package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"InkaTry/warehouse-storage-be/mocks/mock_mysql"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListProducts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		params       *dtos.ListProductsRequest
		expectations func()
		results      func(response *dtos.ListProductsResponse, err error)
	}{
		{
			caseName: "db error",
			params: &dtos.ListProductsRequest{
				CountryID: 1,
			},
			expectations: func() {
				err := errors.New("any")

				mockMysql.EXPECT().ListProducts(ctx, &stores.ListProductsParams{
					CountryID: int64(1),
					Offset:    int64(0),
					Limit:     int64(21),
				}).Return(nil, err)

			},
			results: func(response *dtos.ListProductsResponse, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "no result found",
			params: &dtos.ListProductsRequest{
				CountryID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().ListProducts(ctx, &stores.ListProductsParams{
					CountryID: int64(1),
					Offset:    int64(0),
					Limit:     int64(21),
				}).Return(stores.Products{}, nil)

			},
			results: func(response *dtos.ListProductsResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errs.ErrNoResultFound, err)
			},
		},
		{
			caseName: "successful",
			params: &dtos.ListProductsRequest{
				CountryID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().
					ListProducts(ctx, &stores.ListProductsParams{
						CountryID: int64(1),
						Offset:    int64(0),
						Limit:     int64(21),
					}).
					Return(stores.Products{
						{
							Country: "ID",
						},
					}, nil)

			},
			results: func(response *dtos.ListProductsResponse, err error) {
				assert.Equal(t,
					stores.Products{
						{
							Country: "ID",
						},
					}, response.Products)
			},
		},
		{
			caseName: "successful with pagination",
			params: &dtos.ListProductsRequest{
				CountryID: 1,
				Page:      2,
				Limit:     3,
			},
			expectations: func() {

				mockMysql.EXPECT().
					ListProducts(ctx, &stores.ListProductsParams{
						CountryID: int64(1),
						Offset:    int64((2 - 1) * 3),
						Limit:     int64(3 + 1),
					}).
					Return(stores.Products{
						{
							Country: "ID",
						},
					}, nil)

			},
			results: func(response *dtos.ListProductsResponse, err error) {
				assert.Equal(t,
					stores.Products{
						{
							Country: "ID",
						},
					}, response.Products)
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		tt.expectations()

		tt.results(handler.ListProducts(ctx, tt.params))
	}

}
