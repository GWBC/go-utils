package expiremap

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type ExpireMap[T any] struct {
	cache *gocache.Cache
}

func (e *ExpireMap[T]) New(expiration, cleanupInterval time.Duration) {
	e.cache = gocache.New(expiration, cleanupInterval)
}

func (e *ExpireMap[T]) Set(k string, v T) {
	e.cache.Set(k, v, gocache.DefaultExpiration)
}

func (e *ExpireMap[T]) Get(k string) T {
	v, _ := e.GetX(k)
	return v
}

func (e *ExpireMap[T]) GetWithExpiration(k string) (T, time.Time) {
	v, t, _ := e.GetXWithExpiration(k)
	return v, t
}

func (e *ExpireMap[T]) GetX(k string) (T, bool) {
	var ret T

	v, ok := e.cache.Get(k)
	if !ok {
		return ret, ok
	}

	return v.(T), ok
}

func (e *ExpireMap[T]) GetXWithExpiration(k string) (T, time.Time, bool) {
	var ret T

	v, ex, ok := e.cache.GetWithExpiration(k)
	if !ok {
		return ret, ex, ok
	}

	return v.(T), ex, ok
}

func (e *ExpireMap[T]) Items() map[string]gocache.Item {
	return e.cache.Items()
}

func (e *ExpireMap[T]) Delete(k string) {
	e.cache.Delete(k)
}

func (e *ExpireMap[T]) OnEvicted(f func(string, interface{})) {
	e.cache.OnEvicted(f)
}

func (e *ExpireMap[T]) Flush() {
	e.cache.Flush()
}

func (e *ExpireMap[T]) DeleteExpired() {
	e.cache.DeleteExpired()
}
