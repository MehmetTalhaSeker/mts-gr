package metastore

type MetaStore interface {
	Get(key string) interface{}
	Put(key string, val string) (bool, error)
	Del(key string)
	Flush()
}
