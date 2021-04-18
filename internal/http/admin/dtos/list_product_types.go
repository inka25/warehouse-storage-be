package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListProductTypesResponse struct {
	ProductTypes stores.Results `json:"product_types"`
}
