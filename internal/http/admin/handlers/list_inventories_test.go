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

func TestListInventories(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		params       *dtos.ListInventoriesRequest
		expectations func()
		results      func(response *dtos.ListInventoriesResponse, err error)
	}{
		{
			caseName: "db error",
			params: &dtos.ListInventoriesRequest{
				WarehouseID: 1,
			},
			expectations: func() {
				err := errors.New("any")

				mockMysql.EXPECT().ListInventories(ctx, &stores.ListInventoriesParams{
					WarehouseID: int64(1),
					Offset:      int64(0),
					Limit:       int64(21),
				}).Return(nil, err)

			},
			results: func(response *dtos.ListInventoriesResponse, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "no result found",
			params: &dtos.ListInventoriesRequest{
				WarehouseID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().ListInventories(ctx, &stores.ListInventoriesParams{
					WarehouseID: int64(1),
					Offset:      int64(0),
					Limit:       int64(21),
				}).Return(stores.Inventories{}, nil)

			},
			results: func(response *dtos.ListInventoriesResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errs.ErrNoResultFound, err)
			},
		},
		{
			caseName: "successful",
			params: &dtos.ListInventoriesRequest{
				WarehouseID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().
					ListInventories(ctx, &stores.ListInventoriesParams{
						WarehouseID: int64(1),
						Offset:      int64(0),
						Limit:       int64(21),
					}).
					Return(stores.Inventories{
						{
							Warehouse: "test",
						},
					}, nil)

			},
			results: func(response *dtos.ListInventoriesResponse, err error) {
				assert.Equal(t,
					stores.Inventories{
						{
							Warehouse: "test",
						},
					}, response.Inventories)
			},
		},
		{
			caseName: "successful with pagination",
			params: &dtos.ListInventoriesRequest{
				WarehouseID: 1,
				Page:        2,
				Limit:       3,
			},
			expectations: func() {

				mockMysql.EXPECT().
					ListInventories(ctx, &stores.ListInventoriesParams{
						WarehouseID: int64(1),
						Offset:      int64((2 - 1) * 3),
						Limit:       int64(3 + 1),
					}).
					Return(stores.Inventories{
						{
							Warehouse: "test",
						},
					}, nil)

			},
			results: func(response *dtos.ListInventoriesResponse, err error) {
				assert.Equal(t,
					stores.Inventories{
						{
							Warehouse: "test",
						},
					}, response.Inventories)
			},
		},
	}

	for _, tt := range tts {
		t.Log(tt.caseName)

		tt.expectations()

		tt.results(handler.ListInventories(ctx, tt.params))
	}
}
