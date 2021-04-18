package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListWareshousesResponse struct {
	Warehouses stores.Results `json:"warehouses"`
}
