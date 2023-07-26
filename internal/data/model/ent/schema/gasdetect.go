package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GasDetect holds the schema definition for the GasDetect entity.
type GasDetect struct {
	ent.Schema
}

func (GasDetect) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Post.
func (GasDetect) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("metrics", []map[string]string{}).Optional(),
	}
}
