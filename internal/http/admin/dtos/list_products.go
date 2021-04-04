package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type ListProductsRequest struct{
WarehouseID int64
BrandID int64
ProductTypeID int64
Auto string
Page int64
Limit int64
}

type DownloadListProductRequest ListProductsRequest

type DownloadListProductsResponse struct{
	Filename string
	Products stores.Products
}