package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.String("title").NotEmpty(),
		field.String("short_descr").NotEmpty(),
		field.String("body").NotEmpty(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
