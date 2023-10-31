package parameter

import "time"

type SaleGetParam struct {
	ID string `json:"id,omitempty"`
}

type SaleCreateParam struct {
	ProductID   string `json:"product_id,omitempty"`
	Quantity    int64  `json:"quantity,omitempty"`
	TotalAmount int64  `json:"total_amount,omitempty"`
}

type SaleUpdateParam struct {
	ID          string    `json:"id,omitempty"`
	ProductID   string    `json:"product_id,omitempty"`
	Quantity    int64     `json:"quantity,omitempty"`
	TotalAmount int64     `json:"total_amount,omitempty"`
	Date        time.Time `json:"date,omitempty"`
}

type SaleDeleteParam struct {
	ID string `json:"id,omitempty"`
}
