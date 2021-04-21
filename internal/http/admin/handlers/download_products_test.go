package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"InkaTry/warehouse-storage-be/mocks/mock_mysql"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDownloadProducts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		params       *dtos.DownloadProductRequest
		expectations func()
		results      func(response *dtos.DownloadProductsResponse, err error)
	}{
		{
			caseName: "db error",
			params: &dtos.DownloadProductRequest{
				CountryID: 1,
			},
			expectations: func() {
				err := errors.New("any")

				mockMysql.EXPECT().ListProducts(ctx, &stores.SearchParams{
					CountryID: int64(1),
				}).Return(nil, err)

			},
			results: func(response *dtos.DownloadProductsResponse, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "no result found",
			params: &dtos.DownloadProductRequest{
				CountryID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().ListProducts(ctx, &stores.SearchParams{
					CountryID: int64(1),
				}).Return(stores.Products{}, nil)

			},
			results: func(response *dtos.DownloadProductsResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errs.ErrNoResultFound, err)
			},
		},
		{
			caseName: "successful",
			params: &dtos.DownloadProductRequest{
				CountryID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().
					ListProducts(ctx, &stores.SearchParams{
						CountryID: int64(1),
					}).
					Return(stores.Products{
						{
							Country: "ID",
						},
					}, nil)

			},
			results: func(response *dtos.DownloadProductsResponse, err error) {
				date := time.Now().Format("Monday_02_January_2006")
				expectedFilename := fmt.Sprintf("ProductsList_%s_country_ID", date)
				assert.Equal(t, expectedFilename, response.Filename)
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

		tt.results(handler.DownloadProducts(ctx, tt.params))
	}
}
