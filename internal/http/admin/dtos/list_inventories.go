package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListInventoriesRequest struct {
	WarehouseID int64
	Page        int64
	Limit       int64
}

type ListInventoriesResponse struct {
	Warehouse   string             `json:"warehouse"`
	Inventories stores.Inventories `json:"inventories"`
	HasNext     bool               `json:"has_next"`
	Page        int64              `json:"page"`
}

type DownloadInventoriesRequest ListInventoriesRequest

type DownloadInventoriesResponse struct {
	Filename    string
	Inventories stores.Inventories
}
