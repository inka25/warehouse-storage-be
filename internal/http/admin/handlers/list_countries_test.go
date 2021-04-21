package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"InkaTry/warehouse-storage-be/mocks/mock_mysql"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListCountries(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		expectations func()
		results      func(response *dtos.ListCountriesResponse, err error)
	}{
		{
			caseName: "db error",
			expectations: func() {
				mockMysql.EXPECT().ListCountries(ctx).Return(nil, errors.New("any"))
			},
			results: func(response *dtos.ListCountriesResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errors.New("any"), err)
			},
		},
		{
			caseName: "no result found",
			expectations: func() {
				mockMysql.EXPECT().ListCountries(ctx).Return(stores.Results{}, nil)
			},
			results: func(response *dtos.ListCountriesResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, stores.Results{}, response.Countries)
			},
		},
		{
			caseName: "success",
			expectations: func() {
				mockMysql.EXPECT().ListCountries(ctx).
					Return(stores.Results{
						{
							ID:   uint16(1),
							Name: "test",
						},
					}, nil)
			},
			results: func(response *dtos.ListCountriesResponse, err error) {
				assert.Nil(t, err)
				assert.Equal(t, stores.Results{
					{
						ID:   uint16(1),
						Name: "test",
					},
				}, response.Countries)
			},
		},
	}

	for _, tt := range tts {

		t.Log(tt.caseName)
		tt.expectations()
		tt.results(handler.ListCountries(ctx))
	}
}
