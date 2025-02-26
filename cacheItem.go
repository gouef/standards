package standards

import "time"

type CacheItem interface {
	// GetKey Returns the key for the current cache item.
	GetKey() string
	// Get Retrieves the value of the item from the cache.
	Get() any
	// IsHit Confirms if the cache item lookup resulted in a cache hit.
	IsHit() bool
	// Set Sets the value represented by this cache item
	Set(value any, ttl *time.Duration) (CacheItem, error)
	// ExpiresAt Sets the absolute expiration time for this cache item.
	ExpiresAt(expiration *time.Time) (CacheItem, error)
	// ExpiresAfter Sets the relative expiration time for this cache item.
	ExpiresAfter(t time.Duration)
}
