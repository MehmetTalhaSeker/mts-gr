package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Meta holds the schema definition for the Meta entity.
type Meta struct {
	ent.Schema
}

// Fields of the Meta.
func (Meta) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Unique(),
		field.String("value").NotEmpty(),
	}
}

// Edges of the Meta.
func (Meta) Edges() []ent.Edge {
	return nil
}

func (Meta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
