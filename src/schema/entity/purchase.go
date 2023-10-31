package entity

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	ID          uuid.UUID `json:"id,omitempty"`
	ProductID   uuid.UUID `json:"product_id,omitempty"`
	SupplierID  uuid.UUID `json:"supplier_id,omitempty"`
	Quantity    int64     `json:"quantity,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	TotalAmount int64     `json:"total_amount,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedBy   uuid.UUID `json:"updated_by,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	DeletedBy   uuid.UUID `json:"deleted_by,omitempty"`
	IsDeleted   int64     `json:"is_deleted,omitempty"`
}
