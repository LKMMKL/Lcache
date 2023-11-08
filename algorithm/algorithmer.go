package cache

type Algorithmer interface {
	Get(key string) (interface{}, error)
	Put(key string, value interface{}) error
}
