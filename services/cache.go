package services

import (
    "container/ring"
    "sync"
    "time"
)

const (
    // NumShards is the number of shards in the cache.
    NumShards = 32
    // EvictionTime is the time after which an entry is evicted from the cache.
    EvictionTime = 10 * time.Minute
)

// Cache is a concurrent sharded cache.
type Cache struct {
    shards []*shard
}

// shard is a single shard in the cache.
type shard struct {
    entries *sync.Map
    queue   *ring.Ring
}

// entry is an entry in the cache.
type entry struct {
    value   any
    addedAt time.Time
}

// NewCache creates a new cache.
func NewCache() *Cache {
    c := &Cache{
        shards: make([]*shard, NumShards),
        }
        for i := 0; i < NumShards; i++ {
            c.shards[i] = &shard{
                entries: new(sync.Map),
                queue:   ring.New(1e6), // preallocate queue with 1 million elements
            }
        }
        return c
}

// Set sets the value for the given key in the cache.
func (c *Cache) Set(key string, value any) {
    h := hash(key) % NumShards
    c.shards[h].set(key, value)
}

// Get gets the value for the given key from the cache.
func (c *Cache) Get(key string) (any, bool) {
    h := hash(key) % NumShards
    return c.shards[h].get(key)
}

// set sets the value for the given key in the shard.
func (s *shard) set(key string, value any) {
    s.queue.Value = key
    s.queue = s.queue.Next()
    s.entries.Store(key, &entry{
        value:   value,
        addedAt: time.Now(),
        })
}

// get gets the value for the given key from the shard.
func (s *shard) get(key string) (any, bool) {
    v, ok := s.entries.Load(key)
    if !ok {
        return nil, false
    }
    e, ok := v.(*entry)
    if !ok {
        return nil, false
    }
    if time.Since(e.addedAt) > EvictionTime {
        s.entries.Delete(key)
        return nil, false
    }
    return e.value, true
}

// hash is a simple hash function that returns a non-negative integer for a given string.
func hash(s string) int {
    h := 0
    for _, c := range s {
        h = h*31 + int(c)
    }
    return h
}