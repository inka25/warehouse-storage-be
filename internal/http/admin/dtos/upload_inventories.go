package dtos

type UploadInventoriesRequest struct {
	WarehouseID int64
}

type UploadInventoriesTemplateResponse struct {
	Filename          string
	UploadInventories UploadInventories
}

type UploadInventories []UploadInventory

type UploadInventory struct {
	ID        int64  `json:"product_id" csv:"product_id"`
	Warehouse string `json:"warehouse,omitempty" csv:"warehouse"`
	Code      string `json:"product_code,omitempty" csv:"product_code" `
	Name      string `json:"product_name,omitempty" csv:"product_name" `
	Brand     string `json:"product_brand,omitempty" csv:"product_brand"`
	Type      string `json:"product_type,omitempty" csv:"product_type"`
	Stock     int64  `json:"closing_stock"`
}
