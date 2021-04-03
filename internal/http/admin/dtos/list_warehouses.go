package dtos

type ListWarehousesResponse struct {
	Result []Warehouse `json:"result"`
}

type Warehouse struct {
	ID   uint16 `json:"id"`
	Name string `json:"name"`
}
