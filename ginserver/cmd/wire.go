//go:build wireinject
// +build wireinject

package main

import (
	"github.com/MehmetTalhaSeker/mts-gr/entdb"
	"github.com/MehmetTalhaSeker/mts-gr/ginserver"
	"github.com/MehmetTalhaSeker/mts-gr/metadb"
	"github.com/MehmetTalhaSeker/mts-gr/metastore"
	"github.com/google/wire"
	"net/http"
)

func BuildServerCompileTime() *http.Server {
	wire.Build(
		entdb.NewClient,
		metadb.NewMetaDBMS,
		wire.Bind(new(metastore.MetaStore), new(*metadb.MetaDBMS)),
		ginserver.NewServer,
	)

	return &http.Server{}
}
