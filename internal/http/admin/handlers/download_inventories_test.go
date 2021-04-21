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

func TestDownloadInventories(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMysql := mock_mysql.NewMockClienter(ctrl)

	ctx := context.Background()
	handler := NewAdminHandler(&Params{
		DB: mockMysql,
	})

	var tts = []struct {
		caseName     string
		params       *dtos.DownloadInventoriesRequest
		expectations func()
		results      func(response *dtos.DownloadInventoriesResponse, err error)
	}{
		{
			caseName: "db error",
			params: &dtos.DownloadInventoriesRequest{
				WarehouseID: 1,
			},
			expectations: func() {
				err := errors.New("any")

				mockMysql.EXPECT().ListInventories(ctx, &stores.ListInventoriesParams{
					WarehouseID: int64(1),
				}).Return(nil, err)

			},
			results: func(response *dtos.DownloadInventoriesResponse, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "no result found",
			params: &dtos.DownloadInventoriesRequest{
				WarehouseID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().ListInventories(ctx, &stores.ListInventoriesParams{
					WarehouseID: int64(1),
				}).Return(stores.Inventories{}, nil)

			},
			results: func(response *dtos.DownloadInventoriesResponse, err error) {
				assert.NotNil(t, err)
				assert.Equal(t, errs.ErrNoResultFound, err)
			},
		},
		{
			caseName: "successful",
			params: &dtos.DownloadInventoriesRequest{
				WarehouseID: 1,
			},
			expectations: func() {

				mockMysql.EXPECT().
					ListInventories(ctx, &stores.ListInventoriesParams{
						WarehouseID: int64(1),
					}).
					Return(stores.Inventories{
						{
							Warehouse: "test",
						},
					}, nil)

			},
			results: func(response *dtos.DownloadInventoriesResponse, err error) {
				expectedFilename := fmt.Sprintf("TESTInventory_%s", time.Now().Format("Monday_02_January_2006"))
				assert.Equal(t, expectedFilename, response.Filename)
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

		tt.results(handler.DownloadInventories(ctx, tt.params))
	}

}
