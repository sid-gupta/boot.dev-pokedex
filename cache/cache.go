package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache[T any] interface {
	Add(key string, val T)
	Get(key string) (T, bool)
}

type cacheEntry[T any] struct {
	createdAt time.Time
	entry     T
}

type ExpiringCache[T any] struct {
	lock  sync.RWMutex
	cache map[string]cacheEntry[T]
	ttl   time.Duration
}

func (c *ExpiringCache[T]) Add(key string, val T) {
	createdAt := time.Now()
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[key] = cacheEntry[T]{
		createdAt: createdAt,
		entry:     val,
	}
}

func (c *ExpiringCache[T]) Get(key string) (T, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	val, found := c.cache[key]
	if !found {
		var out T
		return out, false
	}

	expiry := val.createdAt.Add(c.ttl)
	if expiry.Before(time.Now()) {
		var out T
		return out, false
	}

	return val.entry, true
}

func (c *ExpiringCache[T]) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()

	expire := func(now time.Time) {
		c.lock.Lock()
		defer c.lock.Unlock()

		for k, v := range c.cache {
			expiry := v.createdAt.Add(c.ttl)
			if expiry.Before(now) {
				fmt.Printf("\nReaping: %s", k)
				delete(c.cache, k)
			}
		}
	}

	for {
		select {
		case now := <-ticker.C:
			expire(now)
		}
	}
}

func NewExpiringCache[T any](timeout time.Duration) *ExpiringCache[T] {
	cache := &ExpiringCache[T]{
		lock:  sync.RWMutex{},
		cache: make(map[string]cacheEntry[T]),
		ttl:   timeout,
	}

	go cache.reapLoop()

	return cache
}
