package cache

import (
	"testing"
	"time"
)

func TestReapLoop(t *testing.T) {
	const ttl = 2 * time.Second

	expiringCache := NewExpiringCache[string](ttl)

	expiringCache.Add("foo", "barbar")

	_, found := expiringCache.Get("foo")
	if !found {
		t.Errorf("should have found 'foo'")
		return
	}

	time.Sleep(1 * time.Second)

	_, found = expiringCache.Get("foo")
	if !found {
		t.Errorf("should have found 'foo'")
		return
	}

	time.Sleep(2 * time.Second)

	_, found = expiringCache.Get("foo")
	if found {
		t.Errorf("should not have found 'foo'")
		return
	}

	return
}
