package entity

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Country     string    `json:"country,omitempty"`
	Province    string    `json:"province,omitempty"`
	Regency     string    `json:"regency,omitempty"`
	SubDistrict string    `json:"sub_district,omitempty"`
	UserID      uuid.UUID `json:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedBy   uuid.UUID `json:"updated_by,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	DeletedBy   uuid.UUID `json:"deleted_by,omitempty"`
	IsDeleted   int64     `json:"is_deleted,omitempty"`
}
