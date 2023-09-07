package stores

import "time"

type JSONTime string

type Results []Result

type Result struct {
	ID   uint16 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type SearchParams struct {
	ProductID     int64
	WarehouseID   int64
	CountryID     int64
	BrandID       int64
	ProductTypeID int64
	Prefix        string
	Offset        int64
	Limit         int64
}

type UploadParams struct {
	Email       string
	WarehouseId int64
	Timestamp   time.Time
	UploadInfo  Inventories
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
	ID        int64  `json:"id,omitempty" db:"id"`
	Warehouse string `json:"warehouse,omitempty" csv:"-" db:"warehouse"`
	Country   string `json:"country,omitempty" csv:"-" db:"country"`
	Code      string `json:"code,omitempty" db:"code"`
	Name      string `json:"name,omitempty" db:"name"`
	Brand     string `json:"brand,omitempty" db:"brand"`
	Type      string `json:"type,omitempty" db:"type"`
	Stock     int64  `json:"stock,omitempty" db:"stock"`
}

type Histories []History

type History struct {
	EditType    string   `json:"edit_type" db:"edit_type"`
	Description string   `json:"description" db:"description"`
	UpdatedAt   JSONTime `json:"updated_at" db:"updated_at"`
	UpdatedBy   string   `json:"updated_by" db:"updated_by"`
}

func (t *JSONTime) Scan(value interface{}) error {
	timeString := value.(time.Time).Format("Mon, 02 Jan 2006 15:04:05")
	*t = JSONTime(timeString)
	return nil
}
