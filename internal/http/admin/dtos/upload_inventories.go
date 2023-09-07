package dtos

type UploadInventoriesRequest struct {
	Email             string
	IsAdmin           bool
	WarehouseID       int64
	UploadInventories UploadInventories
}

type UploadInventoriesTemplateResponse struct {
	Filename          string
	UploadInventories UploadInventories
}

type UploadInventories []UploadInventory

type UploadInventory struct {
	ID        int64  `json:"ID" csv:"ID"`
	Warehouse string `json:"Warehouse,omitempty" csv:"Warehouse"`
	Code      string `json:"Code,omitempty" csv:"Code" `
	Name      string `json:"Name,omitempty" csv:"Name" `
	Brand     string `json:"Brand,omitempty" csv:"Brand"`
	Type      string `json:"Type,omitempty" csv:"Type"`
	Stock     int64  `json:"Stock" csv:"Stock"`
}
