package core

import "fmt"

// Cache provides an in-memory cache system
type Cache struct {
	data map[string]interface{}
}

// NewCache initializes a new cache instance
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set stores a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	val, found := c.data[key]
	return val, found
}
