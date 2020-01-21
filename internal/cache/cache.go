package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache defines wrapper over the cache
type Cache struct {
	c *cache.Cache
}

// New provides initialization of the cache
func New() *Cache {
	return &Cache{
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

// Set provides inserting to the cache
func (c *Cache) Set(key string, value interface{}) error {
	return c.Set(key, value)
}

// Get provides getting from the cache
func (c *Cache) Get(key string) (interface{}, error) {
	return c.Get(key)
}
