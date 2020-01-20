package cache

import (
	"github.com/patrickmn/go-cache"
)

// Cache defines wrapper over the cache
type Cache struct {
	c *cache.Cache
}

// New provides initialization of the cache
func New()*Cache {
	return &Cache {
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	return c.Set(key, value)
}

func (c *Cache) Get(key string)(interface{}, error) {
	return c.Get(key, value)
}
