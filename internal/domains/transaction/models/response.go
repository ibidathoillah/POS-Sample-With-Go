package models

type OmzetResponse struct {
	MerchantName string `json:"merchant_name" db:"merchant_name"`
	OutletName   string `json:"outlet_name,omitempty" db:"outlet_name"`
	Omzet        string `json:"omzet" db:"omzet"`
	Date         string `json:"date" db:"date"`
}

type TransactionOmzet struct {
	MerchantID   int64   `json:"merchant_id" db:"merchant_id"`
	MerchantName string  `json:"merchant_name" db:"merchant_name"`
	OutletID     *int64  `json:"outlet_id,omitempty" db:"outlet_id"`
	OutletName   *string `json:"outlet_name,omitempty" db:"outlet_name"`
	Omzet        int64   `json:"omzet" db:"omzet"`
	Date         string  `json:"date" db:"date"`
	NumDate      int64   `json:"-" db:"day"`
}
