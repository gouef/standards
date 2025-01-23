package standards

import "time"

type CacheSimple interface {
	// Get Returns a value from the cache.
	Get(key string, defaultValue any) any
	// GetMultiply Returns a list of cache items.
	GetMultiply(keys []string, defaultValue any) []any
	// Has Determines whether an item is present in the cache.
	Has(key string) bool
	// Clear Deletes all cache's keys.
	Clear() error
	// Delete Remove an item from the cache.
	Delete(key string) error
	// DeleteMultiply Removes multiple items in a single operation.
	DeleteMultiply(keys ...string) error
	// Set Persists a cache item.
	Set(key string, item any) error
	// SetMultiply Persists a cache items.
	SetMultiply(values []any, ttl time.Duration) error
}
