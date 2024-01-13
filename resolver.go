package mts_gr

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/MehmetTalhaSeker/mts-gr/ent"
	"github.com/MehmetTalhaSeker/mts-gr/metastore"
)

// Resolver is the resolver root.
type Resolver struct {
	client *ent.Client
	mt     metastore.MetaStore
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, mt metastore.MetaStore) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client, mt},
	})
}
