package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListProductTypesResponse struct {
	ProductTypes stores.ProductTypes `json:"product_types"`
}
