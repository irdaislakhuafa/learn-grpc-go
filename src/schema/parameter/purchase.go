package parameter

import "time"

type PurchaseGetParam struct {
	ID string
}

type PurchaseCreateParam struct {
	ProductID   string    `json:"product_id,omitempty"`
	SupplierID  string    `json:"supplier_id,omitempty"`
	Quantity    string    `json:"quantity,omitempty"`
	TotalAmount int64     `json:"total_amount,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

type PurchaseUpdateParam struct {
	ID          string    `json:"id,omitempty"`
	ProductID   string    `json:"product_id,omitempty"`
	SupplierID  string    `json:"supplier_id,omitempty"`
	Quantity    string    `json:"quantity,omitempty"`
	TotalAmount int64     `json:"total_amount,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

type PurchaseDeleteParam struct {
	ID string `json:"id,omitempty"`
}
