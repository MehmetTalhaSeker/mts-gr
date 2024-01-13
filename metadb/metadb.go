package metadb

import (
	"context"
	"github.com/MehmetTalhaSeker/mts-gr/ent"
	"github.com/MehmetTalhaSeker/mts-gr/ent/meta"
)

type MetaDBMS struct {
	client *ent.Client
}

func NewMetaDBMS(client *ent.Client) *MetaDBMS {
	return &MetaDBMS{
		client: client,
	}
}

func (db *MetaDBMS) Get(key string) interface{} {
	item, err := db.client.Meta.Query().Where(meta.Key(key)).Only(context.Background())
	if err != nil {
		return nil
	}
	return item.Value
}

func (db *MetaDBMS) Put(key string, val string) (bool, error) {
	_, err := db.client.Meta.Create().
		SetKey(key).
		SetValue(val).
		Save(context.Background())

	if err != nil {
		return false, err
	}

	return true, nil
}

func (db *MetaDBMS) Del(key string) {
	item, err := db.client.Meta.Query().Where(meta.Key(key)).Only(context.Background())
	if err != nil {
		return
	}

	_ = db.client.Meta.DeleteOne(item).Exec(context.Background())
}

func (db *MetaDBMS) Flush() {
	_, _ = db.client.Meta.Delete().Exec(context.Background())
}
