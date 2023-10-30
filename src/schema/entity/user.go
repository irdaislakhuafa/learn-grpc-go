package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Age       int       `json:"age,omitempty"`
	Hobbies   []string  `json:"hobbies,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	DeletedBy uuid.UUID `json:"deleted_by,omitempty"`
	Address   Address   `json:"address,omitempty"`
	IsDeleted int64     `json:"is_deleted,omitempty"`
}
