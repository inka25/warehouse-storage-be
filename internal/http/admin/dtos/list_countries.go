package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListCountriesResponse struct {
	Countries stores.Results `json:"countries"`
}
