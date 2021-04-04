package stores

type Warehouses []Warehouse

type Warehouse struct {
	ID   uint16 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Products []Product

type Product struct{
ID int64 `json:"id" db:"id"`
Code string `json:"code" db:"code"`
Name string `json:"name" db:"name"`
Brand string `json:"brand" db:"brand"`
Type string `json:"type" db:"type"`
Stock int64 `json:"stock" db:"stock"`
}