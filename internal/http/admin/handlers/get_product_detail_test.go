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

func TestGetProductDetail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		param        *dtos.GetProductDetailRequest
		expectations func()
		results      func(response *dtos.GetProductDetailResponse, err error)
	}{
		{
			caseName: "db error get product",
			param: &dtos.GetProductDetailRequest{
				ProductId: int64(1),
			},
			expectations: func() {
				mockMysql.EXPECT().GetProductByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(nil, errors.New("err db product"))
			},
			results: func(response *dtos.GetProductDetailResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errors.New("err db product"), err)
			},
		},
		{
			caseName: "db error list histories",
			param: &dtos.GetProductDetailRequest{
				ProductId: int64(1),
			},
			expectations: func() {
				mockMysql.EXPECT().GetProductByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(&stores.Product{}, nil)
				mockMysql.EXPECT().ListHistoriesByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(nil, errors.New("err db histories"))
			},
			results: func(response *dtos.GetProductDetailResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errors.New("err db histories"), err)
			},
		},
		{
			caseName: "db error list inventories",
			param: &dtos.GetProductDetailRequest{
				ProductId: int64(1),
			},
			expectations: func() {
				mockMysql.EXPECT().GetProductByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(&stores.Product{}, nil)
				mockMysql.EXPECT().ListHistoriesByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(stores.Histories{}, nil)
				mockMysql.EXPECT().ListInventoriesByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(nil, errors.New("err db inventories"))
			},
			results: func(response *dtos.GetProductDetailResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errors.New("err db inventories"), err)
			},
		},
		{
			caseName: "no product found",
			param: &dtos.GetProductDetailRequest{
				ProductId: int64(1),
			},
			expectations: func() {
				mockMysql.EXPECT().GetProductByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(nil, errs.ErrNoResultFound)
			},
			results: func(response *dtos.GetProductDetailResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errs.ErrNoResultFound, err)
			},
		},
		{
			caseName: "success",
			param: &dtos.GetProductDetailRequest{
				ProductId: int64(1),
			},
			expectations: func() {
				mockMysql.EXPECT().GetProductByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(&stores.Product{
					ID: int64(1),
				}, nil)
				mockMysql.EXPECT().ListHistoriesByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(stores.Histories{
					{
						UpdatedBy: "admin_test",
					},
				}, nil)
				mockMysql.EXPECT().ListInventoriesByProductId(ctx, &stores.SearchParams{
					ProductID: int64(1),
				}).Return(stores.Inventories{
					{
						Warehouse: "test",
						Stock:     1,
					},
				}, nil)
			},
			results: func(response *dtos.GetProductDetailResponse, err error) {
				assert.Nil(t, err)
				assert.ObjectsAreEqualValues(stores.Product{
					ID: int64(1),
				}, response.Product)
				assert.ObjectsAreEqualValues(stores.Histories{
					{
						UpdatedBy: "admin_test",
					},
				}, response.Histories)
				assert.ObjectsAreEqualValues(stores.Inventories{
					{
						Warehouse: "test",
						Stock:     1,
					},
				}, response.Inventories)
			},
		},
	}

	for _, tt := range tts {

		t.Log(tt.caseName)
		tt.expectations()
		tt.results(handler.GetProductDetail(ctx, tt.param))
	}

}
