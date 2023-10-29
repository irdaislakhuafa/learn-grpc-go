package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	fields := append([]ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.Int("age").Positive(),
		field.Strings("hobbies").Optional().Default([]string{}),
	}, GetBaseFields()...)
	return fields
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
