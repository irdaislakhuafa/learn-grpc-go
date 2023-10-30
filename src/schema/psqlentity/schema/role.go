package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	fields := append([]ent.Field{
		field.String("name").Unique(),
		field.String("description").Optional(),
	}, GetBaseFields()...)
	return fields
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
