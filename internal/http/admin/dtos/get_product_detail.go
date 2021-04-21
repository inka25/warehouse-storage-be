package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type GetProductDetailRequest struct {
	ProductId int64
}

type GetProductDetailResponse struct {
	Product     stores.Product     `json:"product"`
	Inventories stores.Inventories `json:"inventories"`
	Histories   stores.Histories   `json:"histories"`
}
