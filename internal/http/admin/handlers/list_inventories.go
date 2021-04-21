package handlers

import (
	"InkaTry/warehouse-storage-be/internal/http/admin/dtos"
	"InkaTry/warehouse-storage-be/internal/pkg/errs"
	"InkaTry/warehouse-storage-be/internal/pkg/stores"
	"context"
	"log"
)

const logListInventories = "[ListInventories]"

func (h *Handler) ListInventories(ctx context.Context, p *dtos.ListInventoriesRequest) (*dtos.ListInventoriesResponse, error) {
	var res dtos.ListInventoriesResponse
	// check and set defaultValues
	validPageLimit(&p.Page, &p.Limit)

	offset := (p.Page - 1) * p.Limit
	limit := p.Limit + 1

	result, err := h.db.ListInventoriesByWarehouseId(ctx, &stores.SearchParams{
		WarehouseID: p.WarehouseID,
		Offset:      offset,
		Limit:       limit,
	})
	if err != nil {
		log.Printf(
			"%s err: %v\n",
			logListProducts, err)
		return nil, err
	}

	if len(result) == 0 {
		return nil, errs.ErrNoResultFound
	}

	if len(result) > int(p.Limit) {
		res.HasNext = true
		result = result[:p.Limit]
	}
	res.Warehouse = result[0].Warehouse
	res.Page = p.Page
	res.Inventories = result

	return &res, err
}

func validPageLimit(page, limit *int64) {

	if *limit == 0 {
		*limit = 20
	}

	if *page == 0 {
		*page = 1
	}

}
