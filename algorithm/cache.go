package cache

import (
	"errors"
	"math"
	"sync"
)

type Cache struct {
	mutext     sync.Mutex
	maxbytes   int
	algorithme Algorithmer
}

func NewCache(maxbytes int, algorithme Algorithmer) (cache *Cache, err error) {
	if algorithme == nil {

	}
	if maxbytes <= 0 || maxbytes > math.MaxInt32 {
		err = errors.New("maxbytes must be > 0 and < math.MaxInt32")
		return
	}
	cache = &Cache{
		maxbytes:   maxbytes,
		algorithme: algorithme,
	}
	return
}
