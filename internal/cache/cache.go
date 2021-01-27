package cache

import (
	"time"

	"github.com/saromanov/diselfuel/internal/config"
	"github.com/patrickmn/go-cache"
)

// Cache defines wrapper over the cache
type Cache struct {
	c *cache.Cache
}

// New provides initialization of the cache
func New(с *config.Config) *Cache {
	if c == nil {
		panic("config is not defined")
	}
	return &Cache{
		c: cache.New(c.CacheTimeout, c.CacheTimeout),
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
