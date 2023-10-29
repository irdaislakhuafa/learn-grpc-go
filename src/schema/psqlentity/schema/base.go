package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

func GetBaseFields() []ent.Field {
	fields := []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.UUID("created_by", uuid.UUID{}).Unique().Default(uuid.New),
		field.Time("updated_at").Default(time.Now).Optional(),
		field.UUID("updated_by", uuid.UUID{}).Unique().Default(uuid.New).Optional(),
		field.Time("deleted_at").Default(time.Now).Optional(),
		field.UUID("deleted_by", uuid.UUID{}).Unique().Default(uuid.New).Optional(),
		field.Int64("is_deleted").Default(0).Positive(),
	}
	return fields
}
