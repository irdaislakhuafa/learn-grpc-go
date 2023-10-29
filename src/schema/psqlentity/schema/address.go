package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	fields := append([]ent.Field{
		field.String("country").NotEmpty(),
		field.String("province").NotEmpty(),
		field.String("regency").NotEmpty(),
		field.String("sub_district").NotEmpty(),
	}, GetBaseFields()...)
	return fields
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return nil
}
