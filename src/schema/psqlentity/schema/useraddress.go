package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserAddress holds the schema definition for the UserAddress entity.
type UserAddress struct {
	ent.Schema
}

// Fields of the UserAddress.
func (UserAddress) Fields() []ent.Field {
	fields := append([]ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("address_id", uuid.UUID{}),
	}, GetBaseFields()...)
	return fields
}

// Edges of the UserAddress.
func (UserAddress) Edges() []ent.Edge {
	return nil
}
