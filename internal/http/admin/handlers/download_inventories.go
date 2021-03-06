package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

const logDownloadInventories = "[DownloadInventories]"

func (h *Handler) DownloadInventories(ctx context.Context, p *dtos.DownloadInventoriesRequest) (*dtos.DownloadInventoriesResponse, error) {
	var res dtos.DownloadInventoriesResponse

	// for download, no need offset-limit
	result, err := h.db.ListInventoriesByWarehouseId(ctx, &stores.SearchParams{
		WarehouseID: p.WarehouseID,
	})
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logDownloadInventories, err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, errs.ErrNoResultFound
	}

	date := time.Now().Format("Monday_02_January_2006")
	res.Filename = fmt.Sprintf("%sInventory_%s", strings.ToUpper(result[0].Warehouse), date)
	res.Inventories = result

	return &res, err
}
