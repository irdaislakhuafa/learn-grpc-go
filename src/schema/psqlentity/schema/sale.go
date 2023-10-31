package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Sale holds the schema definition for the Sale entity.
type Sale struct {
	ent.Schema
}

// Fields of the Sale.
func (Sale) Fields() []ent.Field {
	fields := append([]ent.Field{
		field.UUID("product_id", uuid.UUID{}),
		field.Int64("quantity"),
		field.Int64("total_amount"),
		field.Time("date"),
	}, GetBaseFields()...)
	return fields
}

// Edges of the Sale.
func (Sale) Edges() []ent.Edge {
	return nil
}
