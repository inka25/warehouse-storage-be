package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"context"
)

const logUploadInventories = "[UploadInventories]"

func (h *Handler) UploadInventories(ctx context.Context, p *dtos.UploadInventoriesRequest) (error, []string) {

	//result, err := h.db.ListInventoriesByWarehouseId(ctx, &stores.SearchParams{
	//	WarehouseID: p.WarehouseID,
	//})
	//if err != nil {
	//	log.Printf(
	//		"%s err: %v\n",
	//		logUploadInventories, err)
	//	return err
	//}
	//
	//if len(result) == 0 {
	//	return errs.ErrNoResultFound
	//}

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
