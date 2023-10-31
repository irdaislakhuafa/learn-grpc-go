package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Purchase holds the schema definition for the Purchase entity.
type Purchase struct {
	ent.Schema
}

// Fields of the Purchase.
func (Purchase) Fields() []ent.Field {
	fields := append([]ent.Field{
		field.UUID("product_id", uuid.UUID{}),
		field.UUID("supplier_id", uuid.UUID{}),
		field.Int64("quantity"),
		field.Time("date"),
		field.Int64("total_amount"),
	}, GetBaseFields()...)
	return fields
}

// Edges of the Purchase.
func (Purchase) Edges() []ent.Edge {
	return nil
}
