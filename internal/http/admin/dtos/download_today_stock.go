package dtos

import "InkaTry/warehouse-storage-be/internal/pkg/stores"

type DownloadTodayStockResponse struct{
	Filename string
	Data stores.Products
}
