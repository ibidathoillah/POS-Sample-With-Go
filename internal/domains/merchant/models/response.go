package models

type Merchant struct {
	ID           int64  `json:"id,omitempty" db:"id"`
	UserID       int64  `json:"user_id" db:"user_id"`
	MerchantName string `json:"merchant_name" db:"merchant_name"`
	MerchantID   int64  `json:"merchant_id" db:"merchant_id"`
}
