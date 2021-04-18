package stores

type Results []Result

type Result struct {
	ID   uint16 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type ListProductsParams struct {
	CountryID     int64
	BrandID       int64
	ProductTypeID int64
	Prefix        string
	Offset        int64
	Limit         int64
}

type ListInventoriesParams struct {
	WarehouseID int64
	Offset      int64
	Limit       int64
}

type Products []Product

type Product struct {
	ID      int64  `json:"id" db:"id"`
	Code    string `json:"code" db:"code"`
	Country string `json:"country" db:"country"`
	Name    string `json:"name" db:"name"`
	Brand   string `json:"brand" db:"brand"`
	Type    string `json:"type" db:"type"`
	Stock   int64  `json:"stock" db:"stock"`
}

type Inventories []Inventory

type Inventory struct {
	ID        int64  `json:"id" db:"id"`
	Warehouse string `json:"-" csv:"-" db:"warehouse"`
	Country   string `json:"country" csv:"-" db:"country"`
	Code      string `json:"code" db:"code"`
	Name      string `json:"name" db:"name"`
	Brand     string `json:"brand" db:"brand"`
	Type      string `json:"type" db:"type"`
	Stock     int64  `json:"stock" db:"stock"`
}
