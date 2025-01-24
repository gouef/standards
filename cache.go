package standards

type Cache interface {
	// GetItem Returns a CacheItem representing the specified key.
	GetItem(key string) CacheItem
	// GetItems Returns a list of cache items.
	GetItems(keys ...string) []CacheItem
	// HasItem Confirms if the cache contains specified cache item.
	HasItem(key string) bool
	// Clear Deletes all items in the pool
	Clear() error
	// DeleteItem Removes the item from the pool.
	DeleteItem(key string) error
	// DeleteItems Removes multiple items from the pool.
	DeleteItems(keys ...string) error
	// Save Persists a cache item immediately.
	Save(item CacheItem) error
	// SaveDeferred Sets a cache item to be persisted later.
	SaveDeferred(item CacheItem) error
	// Commit Persists any deferred cache items.
	Commit() error
}
