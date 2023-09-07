package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"time"
)

const logUploadInventories = "[UploadInventories]"

func (h *Handler) UploadInventories(ctx context.Context, p *dtos.UploadInventoriesRequest) (error, []string) {

	if !p.IsAdmin {
		return errs.ErrAuth, nil
	}

	//for _, inventory := range p.UploadInventories {
	//
	//	invJSON, err := json.MarshalIndent(inventory, "", "  ")
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		continue
	//	}
	//	fmt.Println(string(invJSON))
	//}

	timeNow := time.Now()
	var uploadParams stores.UploadParams
	uploadParams.WarehouseId = p.WarehouseID
	uploadParams.Email = p.Email
	uploadParams.Timestamp = timeNow
	copier.Copy(&uploadParams.UploadInfo, &p.UploadInventories)

	if err := h.db.InsertInventoriesByWarehouseId(ctx, &uploadParams); err != nil {
		fmt.Println(err.Error())
		return err, nil
	}

	return nil, nil
}

func (h *Handler) UploadInventoriesTemplate(ctx context.Context) (*dtos.UploadInventoriesTemplateResponse, error) {

	return &dtos.UploadInventoriesTemplateResponse{
		Filename: "upload_inventories_template",
		UploadInventories: dtos.UploadInventories{
			{
				ID:        123,
				Warehouse: "warehouse name required",
				Code:      "test_123 can be empty",
				Name:      "test_name can be empty",
				Brand:     "test_brand can be empty",
				Type:      "test_type can be empty",
				Stock:     123,
			},
		},
	}, nil
}
