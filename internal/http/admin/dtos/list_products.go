package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListProductsRequest struct {
	CountryID     int64
	BrandID       int64
	ProductTypeID int64
	Prefix        string
	Page          int64
	Limit         int64
}

type ListProductsResponse struct {
	Products stores.Products `json:"products"`
	HasNext  bool            `json:"has_next"`
	Page     int64           `json:"page"`
}

type DownloadProductRequest ListProductsRequest

type DownloadProductsResponse struct {
	Filename string
	Products stores.Products
}
