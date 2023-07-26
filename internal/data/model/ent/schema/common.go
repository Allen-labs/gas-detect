package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type BaseMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Unique().Immutable(),
		field.Time("created_time").
			Immutable().
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime"}),
		field.Bool("is_deleted").Default(false),
	}
}
