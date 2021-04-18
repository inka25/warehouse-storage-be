package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListBrandsResponse struct {
	Brands stores.Results `json:"brands"`
}
