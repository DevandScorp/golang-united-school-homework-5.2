package cache

import "time"

type CacheValue struct {
	value string
	useDeadline bool
	deadline time.Time
}
type Cache struct {
	cacheMap map[string]CacheValue
}

func NewCache() Cache {
	return Cache{
		cacheMap: map[string]CacheValue{},
	}
}

func (cache Cache) Get(key string) (string, bool) {
	cachedValue, ok := cache.cacheMap[key]
	if (!ok) {
		return "", false
	}
	if (cachedValue.useDeadline && cachedValue.deadline.Before(time.Now())) {
		return "", false
	}
	return cachedValue.value, true
}

func (cache Cache) Put(key, value string) {
	cache.cacheMap[key] = CacheValue{
		value: value,
		useDeadline: false,
		deadline: time.Time{},
	}
}

func (cache Cache) Keys() []string {
	keys := []string{}

	for key := range cache.cacheMap {
		if (cache.cacheMap[key].useDeadline && cache.cacheMap[key].deadline.After(time.Now())) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {
	cache.cacheMap[key] = CacheValue{
		value: value,
		useDeadline: true,
		deadline: deadline,
	}
}
