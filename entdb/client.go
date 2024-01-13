package entdb

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/MehmetTalhaSeker/mts-gr/ent"
	"github.com/MehmetTalhaSeker/mts-gr/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewClient() *ent.Client {
	ec, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}

	if err := ec.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	return ec
}
