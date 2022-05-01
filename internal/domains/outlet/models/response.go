package models

type Outlet struct {
	ID         int64  `json:"id,omitempty" db:"id"`
	OutletName string `json:"outlet_name" db:"outlet_name"`
	MerchantID int64  `json:"merchant_id" db:"merchant_id"`
}
